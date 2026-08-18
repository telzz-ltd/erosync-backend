package handler

import (
	"crypto/rand"
	"erosync/internal/dto"
	"erosync/internal/model"
	"erosync/internal/store"
	"erosync/pkg/api"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type RegisterHandler struct {
	store store.Store
}

func NewRegisterHandler(store store.Store) *RegisterHandler {
	return &RegisterHandler{store}
}

func (h *RegisterHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest

	passwordHash, err := h.hashPassword(req.Password)
	if err != nil {
		api.BadRequest(w, err.Error())
		return
	}

	user, err := model.NewUser(rand.Text(), req.Name, req.Email, passwordHash)
	if err != nil {
		api.BadRequest(w, err.Error())
		return
	}

	err = h.store.User.Save(r.Context(), user)
	if err != nil {
		api.BadRequest(w, err.Error())
		return
	}
}

func (*RegisterHandler) hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(b), err
}
