package ports

import (
	"context"
	"erosync/internal/users"
)

type Repository interface {
	Save(ctx context.Context, user *users.User) error
	FindByEmail(email string) (*users.User, error)
}
