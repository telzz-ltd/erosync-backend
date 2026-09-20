package users

import (
	"erosync/internal/lib"
	"erosync/internal/otps"

	"github.com/go-chi/chi/v5"
)

type Module struct {
}

func New(repo Repository, otpService otps.Service, tx lib.Tx) *Module {
	return &Module{}
}

func (m *Module) RegisterRoutes(r chi.Router) {

}
