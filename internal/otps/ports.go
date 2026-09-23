package otps

import (
	"context"
)

type FindOTPParam struct {
	Channel   string
	Recipient string
	Purpose   string
}

type Repository interface {
	Save(context.Context, OTP) error
	FindOne(FindOTPParam) (*OTP, error)
	Delete(context.Context, OTP) error
}
