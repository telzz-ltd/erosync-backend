package otps

import (
	"crypto/rand"
	"erosync/internal/store"
	"errors"
	"fmt"
	"log"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	store *store.Store
}

func NewService(s *store.Store) *Service {
	return &Service{s}
}

type ValidateOTPParam struct {
	Code      string
	Purpose   string
	Channel   string
	Recipient string
}

func (s *Service) Validate(param ValidateOTPParam) error {
	otp, err := s.store.Otp.FindOne(store.FindOTPParam{
		Channel:   param.Channel,
		Recipient: param.Recipient,
		Purpose:   param.Purpose,
	})
	if err != nil {
		return err
	}
	if otp == nil {
		return errors.New("otp not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(otp.CodeHash), []byte(param.Code)); err != nil {
		if otp.Valid() {
			otp.IncreaseAttempt()
			if err := s.store.Otp.Save(otp); err != nil {
				log.Printf("unable to save otp: %v", err)
			}
		}
		return err
	}

	if !otp.Valid() {
		if err := s.store.Otp.Delete(otp); err != nil {
			log.Panicln("unable to delete otp: ", err)
		}
		return errors.New("invalid otp")
	}

	return s.store.Otp.Delete(otp)
}

type CreateOTPParam struct {
	Purpose   string
	Channel   string
	Recipient string
	ExpireMin int
}

func (s *Service) Create(param CreateOTPParam) (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(999999))
	if err != nil {
		return "", err
	}

	code := fmt.Sprintf("%06d", n.Int64())
	codeHash, err := bcrypt.GenerateFromPassword([]byte(code), 12)
	if err != nil {
		return "", err
	}

	otp, err := NewOTP(string(codeHash), param.Recipient, param.Purpose, param.Channel, param.ExpireMin)
	if err != nil {
		return "", err
	}

	err = s.store.Otp.Save(otp)
	if err != nil {
		return "", err
	}

	return code, nil
}
