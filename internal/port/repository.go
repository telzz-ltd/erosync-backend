package port

import (
	"context"
	"erosync/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
}

type FindOTPParam struct {
	Channel   string
	Recipient string
	Purpose   string
}

type OtpRepository interface {
	Save(context.Context, domain.OTP) error
	FindOne(param FindOTPParam) (*domain.OTP, error)
	Delete(context.Context, domain.OTP) error
}
