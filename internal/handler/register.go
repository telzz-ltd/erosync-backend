package handler

import "net/http"

type RegisterHandler struct {
}

func NewRegisterHandler() *RegisterHandler {
	return &RegisterHandler{}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (h *RegisterHandler) Execute() (any, error) {
	return nil, nil
}
