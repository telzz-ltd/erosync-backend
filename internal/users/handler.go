package users

import (
	"erosync/internal/lib/app"
	"erosync/internal/lib/security"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	uc *UseCase
}

func NewHandler(uc *UseCase) *Handler {
	return &Handler{uc}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := app.ShouldBindJSON(r, &req); err != nil {
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
	var req LoginRequest
	if err := app.ShouldBindJSON(r, &req); err != nil {
		app.JSON(w, 400, app.Map{"message": err.Error()})
		return
	}

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
	var req VerifyEmailRequest
	if err := app.ShouldBindJSON(r, &req); err != nil {
		app.Error(w, err)
		return
	}

	err := h.sendEmailVerification.Execute(r.Context(), users_uc.SendEmailVerificationCommand{
		UserID: app.GetValue(r, "userId").(string),
	})
	if err != nil {
		app.JSON(w, 500, app.Map{"message": err.Error()})
		return
	}

	app.JSON(w, 200, app.Map{"message": "verification code sent"})
}

func (h *rHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	var req VerifyEmailRequest
	if err := app.ShouldBindJSON(r, &req); err != nil {
		app.JSON(w, 400, app.Map{"message": err.Error()})
		return
	}

	err := h.verifyEmail.Execute(r.Context(), users_uc.VerifyEmailCommand{
		UserID: app.GetValue(r, "userId").(string),
		Code:   req.OtpCode,
	})
	if err != nil {
		app.JSON(w, 400, app.Map{"message": err.Error()})
		return
	}

	app.JSON(w, 200, app.Map{"message": "success"})
}
