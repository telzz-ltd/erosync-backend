package port

import (
	"context"

	"github.com/telzz/erosync-api/internal/domain"
)

type TxManager interface {
	Execute(ctx context.Context, cb func(ctx context.Context) error) error
}

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	FindByID(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}
