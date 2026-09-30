package handler

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/telzz/erosync-api/internal/domain"
	"github.com/telzz/erosync-api/internal/port"
	"github.com/telzz/erosync-api/internal/service"
	"github.com/telzz/erosync-api/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type RegisterHandler struct {
	tx   port.TxManager
	repo port.UserRepository
	jwt  *service.JwtService
	mail *service.MailService
}

func NewRegisterHandler(
	tx port.TxManager,
	repo port.UserRepository,
	jwt *service.JwtService,
	mail *service.MailService,
) *RegisterHandler {
	return &RegisterHandler{
		tx:   tx,
		repo: repo,
		jwt:  jwt,
		mail: mail,
	}
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z]{2,}(?: [a-zA-Z]{2,}){1,2}$"))),
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
	if err != nil {
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

	user := domain.NewUser(rand.Text(), req.Name, req.Email, string(passwordHash))

	err = h.tx.Execute(r.Context(), func(ctx context.Context) error {
		return h.repo.Save(r.Context(), user)
	})
	if err != nil {
		log.Println("Error saving user", err)
		response.Error(w, 500, err.Error(), nil)
		return
	}

	go func() {
		if err := h.mail.SendWelcomeMail(r.Context(), user); err != nil {
			log.Println("Error sending mail: ", err)
		}
	}()

	accessToken, _ := h.jwt.GenerateToken(r.Context(), user.ID, time.Hour)
	refreshToken, _ := h.jwt.GenerateToken(r.Context(), user.ID, 24*time.Hour)

	response.Success(w, response.Map{
		"user":         user,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
