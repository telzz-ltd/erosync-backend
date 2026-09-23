package application

import (
	"erosync/internal/config"
	"erosync/internal/handler"
	"erosync/internal/port"
)

type Application struct {
	config  *config.Config
	store   *port.Store
	handler *handler.Handler
}

func New(cfg *config.Config, store *port.Store) *Application {
	return &Application{
		config: cfg,
		store:  store,
	}
}
