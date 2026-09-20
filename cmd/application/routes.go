package application

import (
	"erosync/internal/infrastructure/http/handlers"
	"erosync/internal/middleware"
	"erosync/internal/users/usecase"

	"github.com/go-chi/chi/v5"
)

func (a *Application) RegisterRoutes(r chi.Router) {

	//usecases
	loginService := usecase.NewLoginUseCase()
	userHandler := handlers.NewUserHandler()
	r.Post("/auth/register", userHandler.Register(authService))
	r.Post("/auth/login", userHandler.Login(authService))

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)

		r.Post("/verification/email/send-otp", userHandler.SendEmailVerificationCode(store, otpService))
		r.Post("/verification/email/verify", userHandler.VerifyEmail(store, otpService))
	})
}
