package users

import (
	"erosync/internal/middleware"
	"erosync/internal/otps"
	"erosync/internal/platform/ports"

	"github.com/go-chi/chi/v5"
)

type Module struct {
	handler *Handler
}

func New(repo Repository, otpService otps.Service, tx ports.TxExecutor) *Module {
	return &Module{}
}

func (m *Module) RegisterRoutes(r chi.Router) {
	r.Post("/auth/register", m.handler.Register)
	r.Post("/auth/login", m.handler.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Post("/verification/email/send-otp", m.handler.SendEmailVerificationCode)
		r.Post("/verification/email/verify", m.handler.VerifyEmail)
	})
}
