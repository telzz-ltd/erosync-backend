package service

import (
	"context"
	"crypto/rand"
	"erosync/internal/domain"
	"erosync/internal/port"
	"erosync/internal/schema"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailExist = errors.New("email already exist")
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{repo}
}

func (uc *UserService) GetByID(id string) (*domain.User, error) {
	return uc.repo.FindByID(id)
}

func (uc *UserService) GetByEmail(email string) (*domain.User, error) {
	return uc.repo.FindByEmail(email)
}

func (uc *UserService) Create(ctx context.Context, param schema.RegisterRequest) (*domain.User, error) {
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

func (s *UserService) VerifyEmail(ctx context.Context, user domain.User) error {
	user.VerifyEmail()
	return s.repo.Save(ctx, user)
}
