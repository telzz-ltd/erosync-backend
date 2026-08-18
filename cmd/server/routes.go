package server

import (
	"erosync/internal/handler"
	"erosync/internal/infrastructure/postgres"
	"erosync/pkg/api"
	"net/http"
)

func (s *Server) RegisterRoutes() *http.ServeMux {
	store := postgres.New()

	login := handler.NewLoginHandler(store)
	register := handler.NewRegisterHandler(store)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.OK(w, "app running fine")
	})

	mux.HandleFunc("POST /auth/login", login.Handle)
	mux.HandleFunc("POST /auth/register", register.Handle)

	return mux
}
