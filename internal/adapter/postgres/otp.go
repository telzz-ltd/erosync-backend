package postgres

import (
	"context"
	"erosync/internal/domain"
	"erosync/internal/port"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OTPRepository struct {
	db *pgxpool.Pool
}

func NewOTPRepository(db *pgxpool.Pool) *OTPRepository {
	return &OTPRepository{db}
}

func (r *OTPRepository) Save(ctx context.Context, otp domain.OTP) error {
	db := GetExecutor(ctx, r.db)

	_, err := db.Exec(ctx,
		`
			INSERT into otps
				(recipient, code_hash, purpose, channel, attempts, max_attempts, created_at, expires_at)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8)
			ON CONFLICT (recipient, purpose, channel) DO UPDATE SET
				code_hash=EXCLUDED.code_hash,
				attempts=EXCLUDED.attempts,
				max_attempts=EXCLUDED.max_attempts,
				expires_at=EXCLUDED.expires_at
			;
		`,
		otp.Recipient,
		otp.CodeHash,
		otp.Purpose,
		otp.Channel,
		otp.Attempts,
		otp.MaxAttempts,
		otp.CreatedAt,
		otp.ExpiresAt,
	)
	return err
}

func (r *OTPRepository) FindOne(param port.FindOTPParam) (*domain.OTP, error) {
	sql := "SELECT * FROM otps WHERE recipient = $1 AND channel = $2 AND purpose = $3 LIMIT 1;"
	rows, err := r.db.Query(context.Background(), sql, param.Recipient, param.Channel, param.Purpose)
	if err != nil {
		return nil, err
	}

	otp, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.OTP])
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func (r *OTPRepository) Delete(ctx context.Context, otp domain.OTP) error {
	db := GetExecutor(ctx, r.db)

	_, err := db.Exec(ctx,
		"DELETE FROM otps WHERE recipient = $1 AND channel = $2 AND purpose = $3;",
		otp.Recipient,
		otp.Channel,
		otp.Purpose,
	)
	return err
}
