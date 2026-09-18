package usecase

import (
	"context"
	"crypto/rand"
	"erosync/internal/lib/security"
	"erosync/internal/users"
	user_ports "erosync/internal/users/ports"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type CreateAccount struct {
	repo user_ports.Repository
}

func NewCreateAccount(repo user_ports.Repository) *CreateAccount {
	return &CreateAccount{repo}
}

type CreateAccountCommand struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type CreateAccountResponse struct {
	User         users.User `json:"user"`
	AccessToken  string     `json:"accessToken"`
	RefreshToken string     `json:"refreshToken"`
}

func (uc *CreateAccount) Execute(ctx context.Context, cmd CreateAccountCommand) (*CreateAccountResponse, error) {
	if user, _ := uc.repo.FindByEmail(cmd.Email); user != nil {
		return nil, errors.New("email already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), 12)
	if err != nil {
		return nil, err
	}

	user, err := users.NewUser(rand.Text(), cmd.Name, cmd.Email, string(passwordHash))
	if err != nil {
		return nil, err
	}

	err = uc.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	accessToken, err := security.GenerateToken(user.ID, string(user.Role), 30*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshToken, err := security.GenerateToken(user.ID, string(user.Role), 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &CreateAccountResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
