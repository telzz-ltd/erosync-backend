package users

import (
	"context"
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{repo}
}

func (uc *UseCase) GetByID(id string) (*User, error) {
	return uc.repo.FindByID(id)
}

func (uc *UseCase) GetByEmail(email string) (*User, error) {
	return uc.repo.FindByEmail(email)
}

func (uc *UseCase) Create(ctx context.Context, param CreateUserRequest) (*User, error) {
	if user, _ := uc.repo.FindByEmail(param.Email); user != nil {
		return nil, ErrEmailExist
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(param.Password), 12)
	if err != nil {
		return nil, err
	}

	user, err := NewUser(rand.Text(), param.Name, param.Email, string(passwordHash))
	if err != nil {
		return nil, err
	}

	err = uc.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
