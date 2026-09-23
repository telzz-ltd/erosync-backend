package application

import (
	"erosync/internal/middleware"
	"erosync/pkg/response"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CORS)
	r.Use(middleware.Recoverer)

	r.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, 200, response.Map{"message": "app working fine"})
	})

	r.HandleFunc("POST /auth/register", app.Handler.Register)
	r.HandleFunc("POST /auth/login", app.Handler.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate(app.Jwt))

		r.HandleFunc("POST /verification/email/send-otp", app.Handler.SendEmailVerificationCode)
		r.HandleFunc("POST /verification/email/verify", app.Handler.VerifyEmail)

	})

	return r
}
