package port

import (
	"context"
	"erosync/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
}

type FindOTPParam struct {
	Channel   string
	Recipient string
	Purpose   string
}

type OtpRepository interface {
	Save(context.Context, domain.OTP) error
	FindOne(param FindOTPParam) (*domain.OTP, error)
	Delete(ctx context.Context, otp domain.OTP) error
}

type BrandRepository interface {
	Save(context.Context, domain.Brand) error
	FindByID(id string) (*domain.Brand, error)
	Find(param map[string]any) ([]domain.Brand, error)
	ExistByName(name string) bool
	Delete(c context.Context, id string) error
}

// Categories
type BrandCategoryRepository interface {
	Save(context.Context, domain.BrandCategory) error
	Find(map[string]any) ([]domain.BrandCategory, error)
	FindByID(id string) (domain.BrandCategory, error)
	Delete(x context.Context, ids []string) error
	BulkInsert(c context.Context, categories []domain.BrandCategory) error
}
