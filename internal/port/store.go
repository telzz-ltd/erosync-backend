package port

import (
	"context"
)

type TxManager interface {
	Execute(ctx context.Context, cb func(ctx context.Context) error) error
}

type Store struct {
	Tx TxManager

	User          UserRepository
	Otp           OtpRepository
	Brand         BrandRepository
	BrandCategory BrandCategoryRepository
}
