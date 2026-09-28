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
	db    *pgxpool.Pool
	table string
}

func NewBrandCategoryRepository(db *pgxpool.Pool) *BrandCategoryRepository {
	return &BrandCategoryRepository{db, "brand_categories"}
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

	sql := fmt.Sprintf("SELECT * FROM %s", r.table)

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
	rows, err := r.db.Query(context.Background(), fmt.Sprintf("SELECT * FROM %s WHERE id = $1;"), id)
	if err != nil {
		return domain.BrandCategory{}, err
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.BrandCategory])
}

func (r *BrandCategoryRepository) Delete(ctx context.Context, ids []string) error {
	_, err := GetExecutor(ctx, r.db).
		Exec(ctx, fmt.Sprintf("DELETE FROM %s WHERE id IN $1", r.table), ids)
	return err
}
func (r *BrandCategoryRepository) BulkInsert(ctx context.Context, categories []domain.BrandCategory) error {
	db := GetExecutor(ctx, r.db)
	b := &pgx.Batch{}

	for _, cat := range categories {
		sql := fmt.Sprintf(`INSERT INTO %s (id, name, description) VALUES($1, $2, $3);`, r.table)
		b.Queue(sql, cat.ID, cat.Name, cat.Description)
	}

	return db.SendBatch(ctx, b).Close()
}
