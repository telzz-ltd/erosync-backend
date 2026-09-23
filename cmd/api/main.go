package main

import (
	"erosync/internal/adapter/postgres"
	"erosync/internal/application"
	"erosync/internal/config"
)

func main() {
	cfg := config.New()
	store := postgres.New(cfg.DatabaseUrl)

	app := application.New(cfg, store)

	h := app.RegisterRoutes()

	app.Run(cfg.Port, h)
}
