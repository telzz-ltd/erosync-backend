package postgres

import (
	"context"
	"erosync/internal/domain"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BrandCategoryRepository struct {
	db *pgxpool.Pool
}

func NewBrandCategoryRepository(db *pgxpool.Pool) *BrandCategoryRepository {
	return &BrandCategoryRepository{db}
}

func (r *BrandCategoryRepository) Save(ctx context.Context, category domain.BrandCategory) error {
	_, err := GetExecutor(ctx, r.db).Exec(ctx,
		"",
		category.ID, category.Name, category.Description,
	)
	return err
}

func (r *BrandCategoryRepository) Find(param map[string]any) ([]domain.BrandCategory, error) {
	conditions := []string{}
	args := []any{}

	if name, ok := param["name"].(string); ok {
		conditions = append(conditions, "name ILIKE")
		args = append(args, "%"+name+"%")
	}

	if ids, ok := param["ids"].([]string); ok {
		conditions = append(conditions, "id IN")
		args = append(args, ids)
	}

	c2 := []string{}

	for i := range conditions {
		con := fmt.Sprintf("%s $%d", conditions[i], i+1)
		c2 = append(c2, con)
	}

	sql := "SELECT * FROM brand_categories"

	if len(c2) > 0 {
		sql += fmt.Sprintf(" WHERE %s", strings.Join(c2, " AND "))
	}

	fmt.Println("Query: ", sql, args)

	rows, err := r.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.BrandCategory])
}

func (r *BrandCategoryRepository) FindByID(id string) (domain.BrandCategory, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM brand_categories WHERE id = $1;", id)
	if err != nil {
		return domain.BrandCategory{}, err
	}
	return pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.BrandCategory])
}

func (r *BrandCategoryRepository) Delete(ctx context.Context, ids []string) error {
	_, err := GetExecutor(ctx, r.db).
		Exec(ctx, "DELETE FROM brand_categories WHERE id IN $1", ids)
	return err
}
