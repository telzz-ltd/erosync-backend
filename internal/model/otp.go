package model

import (
	"errors"
	"strings"
	"time"
)

var (
	OTPChannelEmail = "EMAIL"
	OTPChannelSMS   = "SMS"

	OTPPurposeVerifyEmail   = "verify_email"
	OTPPurposeResetPassword = "reset_password"
)

type OTP struct {
	Recipient   string    `db:"recipient"`
	CodeHash    string    `db:"code_hash"`
	Purpose     string    `db:"purpose"`
	Channel     string    `db:"channel"`
	Attempts    uint8     `db:"attempts"`
	MaxAttempts uint8     `db:"max_attempts"`
	CreatedAt   time.Time `db:"created_at"`
	ExpiresAt   time.Time `db:"expires_at"`
}

func NewOTP(codeHash, recipient, purpose, channel string, expiryMin int) (*OTP, error) {
	if strings.TrimSpace(codeHash) == "" ||
		strings.TrimSpace(recipient) == "" ||
		strings.TrimSpace(purpose) == "" ||
		strings.TrimSpace(channel) == "" {
		return nil, errors.New("fiels must not be empty")
	}

	otp := OTP{
		Recipient:   recipient,
		CodeHash:    codeHash,
		Purpose:     purpose,
		Channel:     channel,
		Attempts:    0,
		MaxAttempts: 5,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(time.Duration(expiryMin) * time.Minute),
	}

	return &otp, nil
}

func (o *OTP) IncreaseAttempt() {
	o.Attempts += 1
}

func (o *OTP) Valid() bool {
	return o.ExpiresAt.After(time.Now()) && o.Attempts < o.MaxAttempts
}
