package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/telzz/erosync-api/internal/domain"
	"github.com/telzz/erosync-api/internal/port"
	"golang.org/x/crypto/bcrypt"
)

type OtpService struct {
	repo port.OtpRepository
}

func NewOtpService(repo port.OtpRepository) *OtpService {
	return &OtpService{repo}
}

func (s *OtpService) Generate(ctx context.Context, recipient string, channel domain.OTPChannel, purpose domain.OTPPurpose) (string, error) {
	codeInt, err := rand.Int(rand.Reader, big.NewInt(999999))
	if err != nil {
		return "", fmt.Errorf("OTP Code generation: %w", err)
	}

	codeStr := fmt.Sprintf("%0d", codeInt.Int64())
	codeHash, err := bcrypt.GenerateFromPassword([]byte(codeStr), 12)
	if err != nil {
		return "", fmt.Errorf("OTP Code Hash: %w", err)
	}

	otp := domain.NewOTP(recipient, string(codeHash), channel, purpose)

	err = s.repo.Save(ctx, otp)
	if err != nil {
		return "", fmt.Errorf("Save OTP: %w", err)
	}

	return codeStr, nil
}
