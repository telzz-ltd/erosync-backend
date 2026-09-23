package handler

import (
	"erosync/internal/lib/app"
	"erosync/internal/lib/security"
	"erosync/internal/notification"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	uc     *UseCase
	mailer *notification.Mailer
}

func NewHandler(uc *UseCase, m *notification.Mailer) *Handler {
	return &Handler{uc, m}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	req, err := app.ShouldBindJSON[CreateUserRequest](r)
	if err != nil {
		app.JSON(w, 400, app.Map{"message": err.Error()})
		return
	}

	user, err := h.uc.Create(r.Context(), req)
	if err != nil {
		app.JSON(w, 500, app.Map{"message": err.Error()})
		return
	}

	accessToken, _ := security.GenerateToken(user.ID, string(user.Role), 30*time.Minute)
	refreshToken, _ := security.GenerateToken(user.ID, string(user.Role), 24*time.Hour)

	resp := AuthResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	app.JSON(w, 201, resp)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	req, err := app.ShouldBindJSON[LoginRequest](r)
	if err != nil {
		app.JSON(w, 400, app.Map{"message": err.Error()})
		return
	}

	app.JSON(w, 200, app.Map{"message": req})
	return

	user, err := h.uc.GetByEmail(req.Email)
	if err != nil || user == nil {
		app.JSON(w, 500, app.Map{"message": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		app.JSON(w, 500, app.Map{"message": "invalid credentials"})
		return
	}

	accessToken, _ := security.GenerateToken(user.ID, string(user.Role), 30*time.Minute)
	refreshToken, _ := security.GenerateToken(user.ID, string(user.Role), 24*time.Hour)

	resp := AuthResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	app.JSON(w, 200, resp)
}

func (h *Handler) SendEmailVerificationCode(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		app.Error(w, errors.New("unauthenticated"))
		return
	}

	user, err := h.uc.GetByID(userID)
	if err != nil || user == nil {
		if err != nil {
			log.Println(err)
		}
		app.Error(w, errors.New("unauthenticated"))
		return
	}

	// result, err := uc.otps.Generate(ctx, ports.GenerateOTPParam{
	// 	Purpose:   otps.PurposeVerifyEmail,
	// 	Channel:   otps.ChannelEmail,
	// 	Recipient: user.Email,
	// })
	// if err != nil {
	// 	return err
	// }

	h.mailer.Send(r.Context(), notification.SendEmailParams{
		Recipients: []string{user.Email},
		Template:   notification.VerifyEmail,
		Subject:    "Verify your account",
		Args: map[string]any{
			"subject": "verify erosync account",
			"name":    strings.Split(user.Name, " ")[0],
			"code":    123678,
		},
	})

	app.JSON(w, 200, app.Map{"message": "verification code sent"})
}

func (h *Handler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	_, err := app.ShouldBindJSON[VerifyEmailRequest](r)
	if err != nil {
		app.JSON(w, 400, app.Map{"message": err.Error()})
		return
	}

	// user, err := uc.userRepo.FindByID(cmd.UserID)
	// if err != nil {
	// 	return nil
	// }

	// err = uc.otps.Validate(ctx, ports.ValidateOTPParam{
	// 	Code: cmd.Code,
	// 	GenerateOTPParam: ports.GenerateOTPParam{
	// 		Channel:   otps.ChannelEmail,
	// 		Purpose:   otps.PurposeVerifyEmail,
	// 		Recipient: user.Email,
	// 	},
	// })
	// if err != nil {
	// 	return err
	// }

	// if user.EmailVerified() {
	// 	return errors.New("email already verified")
	// }

	// user.VerifyEmail()
	// return uc.userRepo.Save(ctx, user)

	app.JSON(w, 200, app.Map{"message": "success"})
}
