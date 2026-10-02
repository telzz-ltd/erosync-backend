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
	panic("method not implemented")
}

func (r *OtpRepository) FindOne(recipient string, channel domain.OTPChannel, purpose domain.OTPPurpose) (*domain.OTP, error) {
	panic("method not implemented")
}

func (r *OtpRepository) Delete(id string) error {
	panic("method not implemented")
}
