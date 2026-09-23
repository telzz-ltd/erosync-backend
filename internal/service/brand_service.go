package service

import (
	"context"
	"crypto/rand"
	"erosync/internal/domain"
	"erosync/internal/port"
	"erosync/internal/schema"
	"errors"

	"github.com/jackc/pgx/v5"
)

type BrandService struct {
	repo         port.BrandRepository
	categoryRepo port.BrandCategoryRepository
}

func NewBrandService(
	repo port.BrandRepository,
	catRepo port.BrandCategoryRepository,
) *BrandService {
	return &BrandService{repo, catRepo}
}

func (s *BrandService) Create(ctx context.Context, req schema.CreateBrandRequest) (domain.Brand, error) {
	brand, err := domain.NewBrand(rand.Text(), req.Name)
	if err != nil {
		return brand, err
	}

	if s.repo.ExistByName(req.Name) {
		return brand, errors.New("name already exists")
	}

	categories, err := s.categoryRepo.Find(map[string]any{"ids": req.CategoryIds})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return brand, errors.New("invalid categoryId")
		}
		return brand, err
	}

	brand.Categories = categories

	if err = s.repo.Save(ctx, brand); err != nil {
		return brand, err
	}

	return brand, nil
}

func (s *BrandService) GetCategories(params map[string]any) ([]domain.BrandCategory, error) {
	return s.categoryRepo.Find(params)
}
