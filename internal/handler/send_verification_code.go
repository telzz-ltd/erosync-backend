package handler

import (
	"log"
	"net/http"

	"github.com/telzz/erosync-api/internal/domain"
	"github.com/telzz/erosync-api/internal/middleware"
	"github.com/telzz/erosync-api/internal/port"
	"github.com/telzz/erosync-api/internal/service"
	"github.com/telzz/erosync-api/pkg/response"
)

type SendEmailCodeHandler struct {
	tx       port.TxManager
	userRepo port.UserRepository
	otps     *service.OtpService
	mail     *service.MailService
}

func NewSendEmailCodeHandler(
	tx port.TxManager,
	userRepo port.UserRepository,
	otps *service.OtpService,
	mailService *service.MailService,
) *SendEmailCodeHandler {
	return &SendEmailCodeHandler{
		tx:       tx,
		userRepo: userRepo,
		otps:     otps,
		mail:     mailService,
	}
}

func (h *SendEmailCodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.UserIDKey).(string)
	user, err := h.userRepo.FindByID(userId)
	if err != nil {
		response.ServerError(w, err.Error())
		return
	}

	if user == nil {
		response.Unauthenticated(w)
		return
	}

	otpCode, err := h.otps.Generate(
		r.Context(),
		user.Email,
		domain.OTPChannelMail,
		domain.OTPPurposeVerifyEmail,
	)
	if err != nil {
		log.Println("OTP Generate Error:", err)
		response.ServerError(w, "an error occurred")
		return
	}

	go func() {
		err := h.mail.SendVerificationCode(r.Context(), *user, otpCode, 10)
		if err != nil {
			log.Println("Error sending verification mail:", err)
		}
	}()
	response.Success(w, nil)
}
