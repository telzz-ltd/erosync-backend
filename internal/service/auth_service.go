package service

import (
	"context"
	"database/sql"
	"erosync/internal/infrastructure/ports"
	"erosync/internal/lib/notification"
	"erosync/internal/lib/security"
	"erosync/internal/schema"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{userService}
}

func (s *AuthService) Register(ctx context.Context, req schema.CreateAccountRequest) (*schema.AuthResponse, error) {
	user, err := s.userService.Create(ctx, req)
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

	return &schema.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req schema.LoginRequest) (*schema.AuthResponse, error) {
	user, err := s.userService.FindByEmail(req.Email)
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

	return &schema.AuthResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) SendEmailVerificationCode(ctx context.Context, userId string) error {
	user, err := s.userService.FindByID(userId)
	if user == nil {
		return err
	}

	result, err := s.otpService.Generate(ctx, ports.GenerateOTPParam{
		Purpose:   otps.PurposeVerifyEmail,
		Channel:   otps.ChannelEmail,
		Recipient: user.Email,
	})
	if err != nil {
		return err
	}

	go func() {
		if err := notification.SendEmailVerificationMail(*user, result.Code, 10); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
