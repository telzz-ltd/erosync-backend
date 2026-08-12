package handler

import (
	"erosync/internal/pkg"
	"erosync/views"
	"erosync/views/components"
	"net/http"
)

type LoginHandler struct {
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (h *LoginHandler) Page(w http.ResponseWriter, r *http.Request) {
	views.Login().Render(r.Context(), w)
}

type LoginDto struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (h *LoginHandler) Submit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		components.FormErrors([]string{err.Error()}).Render(r.Context(), w)
		return
	}

	var req LoginDto
	err = pkg.ParseForm(r, &req)
	if err != nil {
		components.FormErrors([]string{err.Error()}).Render(r.Context(), w)
		return
	}
}
