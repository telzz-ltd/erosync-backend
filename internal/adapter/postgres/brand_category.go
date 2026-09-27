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
	var (
		where = []string{}
		args  = []any{}
	)

	if name, ok := param["name"].(string); ok {
		args = append(args, "%"+name+"%")
		where = append(where, fmt.Sprintf("name ILIKE $%d", len(args)))
	}

	if ids, ok := param["ids"].([]string); ok {
		args = append(args, ids)
		where = append(where, fmt.Sprintf("id = ANY($%d)", len(args)))
	}

	sql := "SELECT * FROM brand_categories"

	if len(where) > 0 {
		sql += fmt.Sprintf(" WHERE %s", strings.Join(where, " AND "))
	}

	rows, err := r.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.BrandCategory])
}

func (r *BrandCategoryRepository) FindByID(id string) (domain.BrandCategory, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM brand_categories WHERE id = $1;", id)
	if err != nil {
		return domain.BrandCategory{}, err
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.BrandCategory])
}

func (r *BrandCategoryRepository) Delete(ctx context.Context, ids []string) error {
	_, err := GetExecutor(ctx, r.db).
		Exec(ctx, "DELETE FROM brand_categories WHERE id IN $1", ids)
	return err
}
func (r *BrandCategoryRepository) BulkInsert(ctx context.Context, categories []domain.BrandCategory) error {
	_, err := r.db.CopyFrom(ctx,
		pgx.Identifier{"brand_categories"},
		[]string{"id", "name", "description"},
		pgx.CopyFromSlice(len(categories), func(i int) ([]any, error) {
			cat := categories[i]
			return []any{cat.ID, cat.Name, cat.Description}, nil
		}),
	)
	return err
}
