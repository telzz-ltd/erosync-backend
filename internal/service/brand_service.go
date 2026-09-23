package service

import (
	"context"
	"crypto/rand"
	"erosync/internal/domain"
	"erosync/internal/port"
	"erosync/internal/schema"
	"errors"
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

	category, err := s.categoryRepo.FindByID(req.CategoryID)
	if err != nil {
		if errors.Is()
	}

	return brand, nil
}
