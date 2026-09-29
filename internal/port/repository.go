package port

import (
	"context"

	"github.com/telzz/erosync-api/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	FindByID(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}
