package handlers

import (
	"erosync/internal/lib/app"
	users_uc "erosync/internal/users/usecase"
	"log"
	"net/http"
)

type UserHandler struct {
	createAccount         *users_uc.CreateAccount
	login                 *users_uc.Login
	sendEmailVerification *users_uc.SendEmailVerification
	verifyEmail           *users_uc.VerifyEmail
}

func NewUserHandler(
	createAccount *users_uc.CreateAccount,
	login *users_uc.Login,
	sendEmailVerification *users_uc.SendEmailVerification,
	verifyEmail *users_uc.VerifyEmail,
) *UserHandler {
	return &UserHandler{
		createAccount,
		login,
		sendEmailVerification,
		verifyEmail,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if h.createAccount == nil {
		log.Fatal("please pass a non nil usecase")
	}

	var req users_uc.CreateAccountCommand
	if err := app.ShouldBindJSON(r, &req); err != nil {
		app.JSON(w, 400, app.H{"message": err.Error()})
		return
	}

	resp, err := h.createAccount.Execute(r.Context(), req)
	if err != nil {
		app.JSON(w, 500, app.H{"message": err.Error()})
		return
	}

	app.JSON(w, 201, resp)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req users_uc.LoginCommand
	if err := app.ShouldBindJSON(r, &req); err != nil {
		app.JSON(w, 400, app.H{"message": err.Error()})
		return
	}

	resp, err := h.login.Execute(r.Context(), req)
	if err != nil {
		app.JSON(w, 500, app.H{"message": err.Error()})
		return
	}

	app.JSON(w, 201, resp)
}

func (h *UserHandler) SendEmailVerificationCode(w http.ResponseWriter, r *http.Request) {
	err := h.sendEmailVerification.Execute(r.Context(), users_uc.SendEmailVerificationCommand{
		UserID: app.GetValue(r, "userId").(string),
	})
	if err != nil {
		app.JSON(w, 500, app.H{"message": err.Error()})
		return
	}

	app.JSON(w, 200, app.H{"message": "verification code sent"})
}

type VerifyEmailRequest struct {
	OtpCode string `json:"code" validate:"required"`
}

func (h *UserHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	var req VerifyEmailRequest
	if err := app.ShouldBindJSON(r, &req); err != nil {
		app.JSON(w, 400, app.H{"message": err.Error()})
		return
	}

	err := h.verifyEmail.Execute(r.Context(), users_uc.VerifyEmailCommand{
		UserID: app.GetValue(r, "userId").(string),
		Code:   req.OtpCode,
	})
	if err != nil {
		app.JSON(w, 400, app.H{"message": err.Error()})
		return
	}

	app.JSON(w, 200, app.H{"message": "success"})
}
