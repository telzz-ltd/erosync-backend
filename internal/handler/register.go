package handler

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jackc/pgx"
	"github.com/telzz/erosync-api/internal/domain"
	"github.com/telzz/erosync-api/internal/port"
	"github.com/telzz/erosync-api/internal/service"
	"github.com/telzz/erosync-api/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type RegisterHandler struct {
	repo port.UserRepository
	jwt  *service.JwtService
}

func NewRegisterHandler(
	repo port.UserRepository,
	jwt *service.JwtService,
) *RegisterHandler {
	return &RegisterHandler{
		repo: repo,
		jwt:  jwt,
	}
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z]{2,}(?: [a-zA-Z]{2,}){2,3}$"))),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(8, 50)),
	)
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, 400, response.MsgInvalidBody, nil)
		return
	}

	if err := req.Validate(); err != nil {
		if vErr, ok := errors.AsType[validation.Errors](err); ok {
			response.Error(w, 400, response.MsgInvalidBody, vErr)
			return
		}
		response.Error(w, 500, err.Error(), nil)
		return
	}

	existingUser, err := h.repo.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		response.Error(w, 500, err.Error(), nil)
		return
	}

	if existingUser != nil {
		response.Error(w, 400, response.MsgInvalidBody, response.Map{"email": "email already exist"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		response.Error(w, 500, err.Error(), nil)
		return
	}

	user := domain.User{
		ID:           rand.Text(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(passwordHash),
		Status:       domain.UserStatusActive,
		Role:         domain.UserRoleRegular,
	}

	err = h.repo.Save(r.Context(), &user)
	if err != nil {
		response.Error(w, 500, err.Error(), nil)
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

func (h *RegisterHandler) Execute() (any, error) {
	return nil, nil
}
