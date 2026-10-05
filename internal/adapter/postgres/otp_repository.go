package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/telzz/erosync-api/internal/domain"
)

type OtpRepository struct {
	db *pgxpool.Pool
}

func NewOtpRepository(db *pgxpool.Pool) *OtpRepository {
	return &OtpRepository{db}
}

func (r *OtpRepository) Save(ctx context.Context, otp domain.OTP) error {
	_, err := GetExecutor(ctx, r.db).Exec(ctx,
		`INSERT INTO otps (code_hash, recipient, channel, purpose, attempts, max_attempts, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (recipient, channel, purpose) DO UPDATE SET
			attempts=EXCLUDED.attempts;`,
		otp.CodeHash,
		otp.Recipient,
		otp.Channel,
		otp.Purpose,
		otp.Attempts,
		otp.MaxAttempts,
		otp.ExpiresAt,
	)
	return err
}

func (r *OtpRepository) FindOne(recipient string, channel domain.OTPChannel, purpose domain.OTPPurpose) (*domain.OTP, error) {
	var otp domain.OTP

	err := r.db.QueryRow(context.Background(),
		`SELECT (code_hash, recipient, channel, purpose, attempts, max_attempts, expires_at) FROM otps 
		WHERE (recipient = $1 AND channel = $2 AND purpose = $3) LIMIT 1;`,
		recipient, channel, purpose,
	).Scan(
		&otp.CodeHash, &otp.Recipient, &otp.Channel, &otp.Purpose, &otp.Attempts,
		&otp.MaxAttempts, &otp.ExpiresAt,
	)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func (r *OtpRepository) Delete(ctx context.Context, otp domain.OTP) error {
	_, err := r.db.Exec(ctx,
		"DELETE FROM otps WHERE recipient = $1 AND channel = $2 AND purpose = $3;",
		otp.Recipient, otp.Channel, otp.Purpose,
	)
	return err
}
