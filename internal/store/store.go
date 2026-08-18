package store

import (
	"context"
	"erosync/internal/model"
)

type UserStore interface {
	Save(ctx context.Context, user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type Store struct {
	User UserStore
}
