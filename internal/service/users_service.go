package service

import (
	"context"
	"crypto/rand"
	"erosync/internal/domain"
	"erosync/internal/ports"
	"erosync/internal/schema"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{}
}

func (s *UserService) Create(ctx context.Context, req schema.CreateAccountRequest) (user domain.User, err error) {
	if existingUser, _ := s.repo.FindByEmail(req.Email); existingUser != nil {
		return user, errors.New("email already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return user, err
	}

	user, err = domain.NewUser(rand.Text(), req.Name, req.Email, string(passwordHash))
	if err != nil {
		return user, err
	}

	err = s.repo.Save(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserService) FindByEmail(email string) (*domain.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) FindByID(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}
