package service

import (
	"context"
	"crypto/rand"
	"erosync/internal/domain"
	"erosync/internal/port"
	"errors"
	"fmt"
	"log"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

type OtpService struct {
	repo port.OtpRepository
}

func NewOtpService(r port.OtpRepository) *OtpService {
	return &OtpService{r}
}

type ValidateOTPParam struct {
	Code      string
	Purpose   string
	Channel   string
	Recipient string
}

func (s *OtpService) Validate(ctx context.Context, param ValidateOTPParam) error {
	otp, err := s.repo.FindOne(port.FindOTPParam{
		Channel:   param.Channel,
		Recipient: param.Recipient,
		Purpose:   param.Purpose,
	})
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(otp.CodeHash), []byte(param.Code)); err != nil {
		if otp.Valid() {
			otp.IncreaseAttempt()
			if err := s.repo.Save(ctx, *otp); err != nil {
				log.Printf("unable to save otp: %v", err)
			}
		}
		return err
	}

	if !otp.Valid() {
		if err := s.repo.Delete(ctx, *otp); err != nil {
			log.Panicln("unable to delete otp: ", err)
		}
		return errors.New("invalid otp")
	}

	return s.repo.Delete(ctx, *otp)
}

type CreateOTPParam struct {
	Purpose   string
	Channel   string
	Recipient string
	ExpireMin int
}

func (s *OtpService) Create(ctx context.Context, param CreateOTPParam) (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(999999))
	if err != nil {
		return "", err
	}

	code := fmt.Sprintf("%06d", n.Int64())
	codeHash, err := bcrypt.GenerateFromPassword([]byte(code), 12)
	if err != nil {
		return "", err
	}

	otp, err := domain.NewOTP(string(codeHash), param.Recipient, param.Purpose, param.Channel, param.ExpireMin)
	if err != nil {
		return "", err
	}

	err = s.repo.Save(ctx, *otp)
	if err != nil {
		return "", err
	}

	return code, nil
}
