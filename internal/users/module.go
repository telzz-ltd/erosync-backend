package users

import (
	"erosync/internal/lib"
	"erosync/internal/middleware"
	"erosync/internal/otps"

	"github.com/go-chi/chi/v5"
)

type Module struct {
}

func New(repo Repository, otpService otps.Service, tx lib.Tx) *Module {
	return &Module{}
}

func (m *Module) RegisterRoutes(r chi.Router) {
	r.Post("/auth/register", handler.Register(authService))
	r.Post("/auth/login", handler.Login(authService))

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)

		r.Post("/verification/email/send-otp", handler.SendEmailVerificationCode(store, otpService))
		r.Post("/verification/email/verify", handler.VerifyEmail(store, otpService))
	})
}
