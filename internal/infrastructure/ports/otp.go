package ports

import (
	"context"
	"time"
)

type GenerateOTPParam struct {
	Recipient string
	Channel   string
	Purpose   string
	ExpiresAt *time.Duration
}

type GenerateOTPResult struct {
	Code string
}

type ValidateOTPParam struct {
	GenerateOTPParam
	Code string
}

type OtpService interface {
	Generate(ctx context.Context, param GenerateOTPParam) (*GenerateOTPResult, error)
	Validate(ctx context.Context, param ValidateOTPParam) error
}
