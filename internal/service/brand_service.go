package service

import (
	"context"
	"crypto/rand"
	"erosync/internal/domain"
	"erosync/internal/port"
	"erosync/internal/schema"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
)

type BrandService struct {
	repo         port.BrandRepository
	categoryRepo port.BrandCategoryRepository
	tx           port.TxManager
}

func NewBrandService(
	repo port.BrandRepository,
	catRepo port.BrandCategoryRepository,
	tx port.TxManager,
) *BrandService {
	return &BrandService{repo, catRepo, tx}
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

	if len(categories) != len(req.CategoryIds) {
		return brand, errors.New("invalid category ids")
	}

	brand.Categories = categories

	err = s.repo.Save(ctx, brand)
	if err != nil {
		log.Println(err)
		return brand, err
	}

	return brand, nil
}

func (s *BrandService) Query(params map[string]any) ([]domain.Brand, error) {
	return s.repo.Find(params)
}

func (s *BrandService) QueryCategories(params map[string]any) ([]domain.BrandCategory, error) {
	return s.categoryRepo.Find(params)
}

func (s *BrandService) CreateCategories(ctx context.Context, req schema.BulkCreateBrandCategoriesRequest) error {
	categories := []domain.BrandCategory{}

	for _, cat := range req.Data {
		categories = append(categories, domain.BrandCategory{
			ID:          rand.Text(),
			Name:        cat.Name,
			Description: &cat.Description,
		})
	}

	return s.categoryRepo.BulkInsert(ctx, categories)
}
