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

func New(
	validator *validator.Validator,
	users *service.UserService,
	otps *service.OtpService,
	mail *service.MailService,
	jwt *service.JwtService,

) *Handler {
	return &Handler{
		validator: validator,
		users:     users,
		otps:      otps,
		mail:      mail,
		jwt:       jwt,
	}
}
