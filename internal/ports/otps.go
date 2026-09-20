package ports

import (
	"context"
	"erosync/internal/domain"
)

type FindOTPParam struct {
	Channel   string
	Recipient string
	Purpose   string
}

type OTPRepository interface {
	Save(context.Context, domain.OTP) error
	FindOne(FindOTPParam) (*domain.OTP, error)
	Delete(context.Context, domain.OTP) error
}
