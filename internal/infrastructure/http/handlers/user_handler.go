package handlers

import (
	"erosync/internal/lib/app"
	users_uc "erosync/internal/users/usecase"
	"log"
	"net/http"
)

type UserHandler struct {
	createAccount *users_uc.CreateAccount
	login         *users_uc.Login
}

func NewUserHandler(
	createAccount *users_uc.CreateAccount,
) *UserHandler {
	return &UserHandler{
		createAccount: createAccount,
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
