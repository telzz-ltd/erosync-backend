package usecase

import (
	"context"
	"database/sql"
	"erosync/internal/lib/security"
	"erosync/internal/users"
	"erosync/internal/users/ports"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	repo ports.Repository
}

type LoginCommand struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type LoginResponse struct {
	User         users.User `json:"user"`
	AccessToken  string     `json:"accessToken"`
	RefreshToken string     `json:"refreshToken"`
}

func (uc *Login) Execute(ctx context.Context, cmd LoginCommand) (*LoginResponse, error) {
	user, err := uc.repo.FindByEmail(cmd.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(cmd.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := security.GenerateToken(user.ID, string(user.Role), 30*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshToken, err := security.GenerateToken(user.ID, string(user.Role), 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
