package handler

import (
	"net/http"

	"github.com/telzz/erosync-api/internal/port"
)

type SendEmailCodeHandler struct {
	tx       port.TxManager
	userRepo port.UserRepository
	otpRepo  port.OtpRepository
}

func NewSendEmailCodeHandler(
	tx port.TxManager,
	userRepo port.UserRepository,
	otpRepo port.OtpRepository,
) *SendEmailCodeHandler {
	return &SendEmailCodeHandler{
		tx:       tx,
		userRepo: userRepo,
		otpRepo:  otpRepo,
	}
}

func (h *SendEmailCodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
