package router

import (
	"erosync/internal/application"
	"erosync/internal/handler"
	"erosync/internal/middleware"
	"erosync/pkg/response"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterV1Router(app *application.Application, h *handler.Handler) chi.Router {
	r := chi.NewRouter()

	r.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, 200, response.Map{"message": "app working fine"})
	})

	r.HandleFunc("POST /auth/register", h.Register)
	r.HandleFunc("POST /auth/login", h.Login)

	r.HandleFunc("GET /brands/categories", h.GetBrandCategories)

	//Authenticated routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticate(app.Jwt))

		r.HandleFunc("POST /verification/email/send-otp", h.SendEmailVerificationCode)
		r.HandleFunc("POST /verification/email/verify", h.VerifyEmail)

		//Admin routes
		r.Group(func(r chi.Router) {
			r.Use(middleware.Admin)

			r.HandleFunc("POST /brands", h.CreateBrand)
		})
	})

	return r
}
