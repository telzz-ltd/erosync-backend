package users

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
}
