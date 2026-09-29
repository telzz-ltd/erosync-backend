package usecase

import (
	"context"
	"crypto/rand"
	"erosync/internal/domain"
	"erosync/internal/port"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailExist = errors.New("email already exists")
)

type CreateUser struct {
	repo port.UserRepository
	tx   port.TxManager
}

func NewCreateUser(repo port.UserRepository, tx port.TxManager) *CreateUser {
	return &CreateUser{repo, tx}
}

type CreateUserParam struct {
	Name     string
	Email    string
	Password string
}

type CreateUserResponse struct {
	User *domain.User
}

func (uc *CreateUser) Execute(ctx context.Context, param CreateUserParam) (*domain.User, error) {
	if user, _ := uc.repo.FindByEmail(param.Email); user != nil {
		return nil, ErrEmailExist
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(param.Password), 12)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(rand.Text(), param.Name, param.Email, string(passwordHash))
	if err != nil {
		return nil, err
	}

	err = uc.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
