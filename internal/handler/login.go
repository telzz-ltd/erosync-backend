package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"

	"github.com/telzz/erosync-api/internal/port"
	"github.com/telzz/erosync-api/internal/service"
	"github.com/telzz/erosync-api/pkg/response"
)

type LoginHandler struct {
	userRepo port.UserRepository
	jwt      *service.JwtService
}

func NewLoginHandler(userRepo port.UserRepository, jwt *service.JwtService) *LoginHandler {
	return &LoginHandler{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required),
	)
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, err.Error(), nil)
		return
	}

	if err := req.Validate(); err != nil {
		if vErr, ok := errors.AsType[validation.Errors](err); ok {
			response.Error(w, 400, "validation error", vErr)
			return
		}

		response.Error(w, 500, err.Error(), nil)
		return
	}

	user, err := h.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.Error(w, 400, "invalid email or password", nil)
			return
		}
		response.Error(w, 500, err.Error(), nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		response.Error(w, 400, "invalid email or password", nil)
		return
	}

	accessToken, _ := h.jwt.GenerateToken(r.Context(), user.ID, time.Hour)
	refreshToken, _ := h.jwt.GenerateToken(r.Context(), user.ID, 24*time.Hour)

	response.Success(w, response.Map{
		"user":         user,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
