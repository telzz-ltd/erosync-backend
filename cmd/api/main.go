package main

import (
	"erosync/internal/adapter/postgres"
	"erosync/internal/application"
	"erosync/internal/config"
	"erosync/internal/handler"
	"erosync/internal/middleware"
	"erosync/internal/router"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CORS)
	r.Use(middleware.Recoverer)

	cfg := config.New()
	store := postgres.New(cfg.DatabaseUrl)

	app := application.New(cfg, store)

	h := handler.New(app)

	r.Mount("/", router.RegisterV1Router(app, h))

	app.Run(cfg.Port, r)
}
