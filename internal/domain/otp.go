package domain

import (
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
