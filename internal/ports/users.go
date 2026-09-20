package ports

import (
	"context"
	"erosync/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
}
