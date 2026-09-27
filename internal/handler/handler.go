package handler

import "erosync/internal/application"

type Handler struct {
	app *application.Application
}

func New(app *application.Application) *Handler {
	return &Handler{app}
}
