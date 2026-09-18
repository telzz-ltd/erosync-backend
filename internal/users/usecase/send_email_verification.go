package usecase

import (
	"context"
	"erosync/internal/infrastructure/ports"
	"erosync/internal/lib/notification"
	"erosync/internal/otps"
	"erosync/internal/users"
	"log"
)

type SendEmailVerification struct {
	repo users.Repository
	otps ports.OtpService
}

type SendEmailVerificationCommand struct {
	UserID string
}

func (uc *SendEmailVerification) Execute(ctx context.Context, cmd SendEmailVerificationCommand) error {
	user, err := uc.repo.FindByID(cmd.UserID)
	if user == nil {
		return err
	}

	result, err := uc.otps.Generate(ctx, ports.GenerateOTPParam{
		Purpose:   otps.PurposeVerifyEmail,
		Channel:   otps.ChannelEmail,
		Recipient: user.Email,
	})
	if err != nil {
		return err
	}

	go func() {
		if err := notification.SendEmailVerificationMail(*user, result.Code, 10); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
