package server

import (
	"erosync/internal/handler"
	"erosync/views"
	"net/http"
)

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	login := handler.NewLoginHandler()
	register := handler.NewRegisterHandler()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if err := views.Home().Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<p>Server Healthy</p>`))
	})

	mux.HandleFunc("GET /login", login.Page)
	mux.HandleFunc("POST /login", login.Submit)

	mux.HandleFunc("GET /register", register.Page)
	mux.HandleFunc("POST /register", register.Submit)
}
