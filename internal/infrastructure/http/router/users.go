package router

import (
	"erosync/internal/infrastructure/http/handlers"
	"erosync/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func (r *Router) RegisterUserRoutes() {
	h := handlers.NewUserHandler()
	r.Post("/auth/register", h.Register(authService))
	r.Post("/auth/login", h.Login(authService))

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)

		r.Post("/verification/email/send-otp", h.SendEmailVerificationCode(store, otpService))
		r.Post("/verification/email/verify", h.VerifyEmail(store, otpService))
	})
}
