package handler

import (
	"erosync/views"
	"net/http"
)

type RegisterHandler struct {
}

func NewRegisterHandler() *RegisterHandler {
	return &RegisterHandler{}
}

func (h *RegisterHandler) Page(w http.ResponseWriter, r *http.Request) {
	views.Register().Render(r.Context(), w)
}

func (h *RegisterHandler) Submit(w http.ResponseWriter, r *http.Request) {
	views.Register().Render(r.Context(), w)
}
