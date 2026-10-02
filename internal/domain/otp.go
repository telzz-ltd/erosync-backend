package domain

import (
	"crypto/rand"
	"time"
)

type OTPChannel string
type OTPPurpose string

var (
	OTPChannelMail OTPChannel = "MAIL"
	OTPChannelSMS  OTPChannel = "SMS"

	OTPPurposeVerifyEmail   OTPPurpose = "VERIFY_EMAIL"
	OTPPurposeResetPassword OTPPurpose = "RESET_PASSWORD"
)

type OTP struct {
	ID          string
	Recipient   string
	CodeHash    string
	Channel     OTPChannel
	Purpose     OTPPurpose
	ExpiresAt   time.Time
	Attempts    int
	MaxAttempts int
}

func NewOTP(recipient, codeHash string, channel OTPChannel, purpose OTPPurpose) OTP {
	return OTP{
		ID:          rand.Text(),
		CodeHash:    codeHash,
		Recipient:   recipient,
		Channel:     channel,
		Purpose:     purpose,
		Attempts:    0,
		MaxAttempts: 5,
		ExpiresAt:   time.Now().Add(10 * time.Minute),
	}
}

func (otp *OTP) GetExpiryMinute() int {
	return int(time.Until(otp.ExpiresAt).Minutes())
}
