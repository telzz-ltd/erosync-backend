package usecase

import (
	"context"
	"erosync/internal/infrastructure/ports"
	"erosync/internal/otps"
	"erosync/internal/users"
	"errors"
)

type VerifyEmail struct {
	userRepo users.Repository
	otps     ports.OtpService
}

type VerifyEmailCommand struct {
	UserID string
	Code   string
}

func (uc *VerifyEmail) Execute(ctx context.Context, cmd VerifyEmailCommand) error {
	user, err := uc.userRepo.FindByID(cmd.UserID)
	if err != nil {
		return nil
	}

	err = uc.otps.Validate(ctx, ports.ValidateOTPParam{
		Code: cmd.Code,
		GenerateOTPParam: ports.GenerateOTPParam{
			Channel:   otps.ChannelEmail,
			Purpose:   otps.PurposeVerifyEmail,
			Recipient: user.Email,
		},
	})
	if err != nil {
		return err
	}

	if user.EmailVerified() {
		return errors.New("email already verified")
	}

	user.VerifyEmail()
	return uc.userRepo.Save(ctx, user)
}
