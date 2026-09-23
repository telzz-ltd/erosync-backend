package handler

import "net/http"

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req 
}
