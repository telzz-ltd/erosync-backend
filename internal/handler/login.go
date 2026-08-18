package handler

import (
	"encoding/json"
	"erosync/internal/dto"
	"erosync/internal/store"
	"erosync/pkg/api"
	"errors"
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	store store.Store
}

func NewLoginHandler(store store.Store) *LoginHandler {
	return &LoginHandler{store}
}

func (h *LoginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.BadRequest(w, err.Error())
		return
	}

	err = req.Validate()
	if err != nil {
		if vErr, ok := errors.AsType[validation.Errors](err); ok {
			api.ValidationError(w, vErr)
			return
		}

		api.BadRequest(w, err.Error())
		return
	}

	user, err := h.store.User.FindByEmail(req.Email)
	if err != nil {
		fmt.Println("err", err.Error())
		api.BadRequest(w, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		api.BadRequest(w, err.Error())
		return
	}

	api.OK(w, user)
}
