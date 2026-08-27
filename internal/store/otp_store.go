package store

import (
	"erosync/internal/model"
	"errors"

	"github.com/jmoiron/sqlx"
)

type otps struct {
	db *sqlx.DB
}

func (u *otps) Save(otp *model.OTP) error {
	sql := `
		INSERT into otps 
			(recipient, code_hash, purpose, channel, attempts, max_attempts, created_at, expires_at)
		VALUES 
			(:recipient, :code_hash, :purpose, :channel, :attempts, :max_attempts, :created_at, :expires_at)
		ON CONFLICT (recipient, purpose, channel) DO UPDATE SET
			code_hash=EXCLUDED.code_hash,
			attempts=EXCLUDED.attempts,
			max_attempts=EXCLUDED.max_attempts,
			expires_at=EXCLUDED.expires_at
		;
	`
	_, err := u.db.NamedExec(sql, otp)
	return err
}

type FindOTPParam struct {
	Channel   string
	Recipient string
	Purpose   string
}

func (o *otps) FindOne(param FindOTPParam) (*model.OTP, error) {
	var otp model.OTP

	sql := "SELECT * FROM otps WHERE recipient = $1 AND channel = $2 AND purpose = $3 LIMIT 1;"
	err := o.db.Get(&otp, sql, param.Recipient, param.Channel, param.Purpose)
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (o *otps) Delete(otp *model.OTP) error {
	if otp == nil {
		return errors.New("otp cannot be nil")
	}
	_, err := o.db.NamedExec("DELETE FROM otps WHERE recipient = :recipient AND channel = :channel AND purpose = :purpose;", otp)
	return err
}
