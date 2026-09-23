package handler

import (
	"erosync/internal/service"
	"erosync/pkg/validator"
)

type Handler struct {
	//utils
	validator *validator.Validator

	//services
	users *service.UserService
	otps  *service.OtpService
	mail  *service.MailService
	jwt   *service.JwtService
}

func New() *Handler {
	return &Handler{}
}
