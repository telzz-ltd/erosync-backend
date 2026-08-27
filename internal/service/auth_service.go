package service

import (
	"crypto/rand"
	"database/sql"
	"erosync/internal/dto"
	"erosync/internal/model"
	"erosync/internal/shared/security"
	"erosync/internal/store"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	store *store.Store
}

func NewAuthService(store *store.Store) *AuthService {
	return &AuthService{store}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {
	if user, _ := s.store.User.FindByEmail(req.Email); user != nil {
		return nil, errors.New("email already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return nil, err
	}

	user, err := model.NewUser(rand.Text(), req.Name, req.Email, string(passwordHash))
	if err != nil {
		return nil, err
	}

	err = s.store.User.Save(user)
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

	return &dto.AuthResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.store.User.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
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

	return &dto.AuthResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
