package postgres

import (
	"context"
	"erosync/internal/domain"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BrandRepository struct {
	db *pgxpool.Pool
}

func NewBrandRepository(db *pgxpool.Pool) *BrandRepository {
	return &BrandRepository{db}
}

func (r *BrandRepository) Save(ctx context.Context, brand domain.Brand) error {
	db := GetExecutor(ctx, r.db)

	batch := &pgx.Batch{}

	batch.Queue(`
		INSERT INTO brands (id, name, description, logo_url, contact_info, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET
			name=EXCLUDED.name,
			description=EXCLUDED.description,
			logo_url=EXCLUDED.logo_url,
			contact_info=EXCLUDED.contact_info
		;
	`,
		brand.ID, brand.Name, brand.Description, brand.LogoUrl, brand.ContactInfo, brand.CreatedAt,
	)

	for _, cat := range brand.Categories {
		batch.Queue(`
			INSERT INTO brand_category_pivot (brand_id, category_id) VALUES($1, $2)
			ON CONFLICT (brand_id, category_id) DO NOTHING;
			`,
			brand.ID, cat.ID,
		)
	}

	return db.SendBatch(ctx, batch).Close()
}

func (r *BrandRepository) FindByID(id string) (*domain.Brand, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM brands WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	brand, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Brand])
	if err != nil {
		return nil, err
	}

	rows, err = r.db.Query(
		context.Background(),
		`SELECT * FROM brand_categories as c 
		INNER JOIN brand_category_pivot as p 
		ON c.id = p.category_id AND p.brand_id = $1`,
		brand.ID,
	)

	brand.Categories, err = pgx.CollectRows(rows, pgx.RowToStructByName[domain.BrandCategory])
	if err != nil {
		return nil, err
	}

	return &brand, nil
}

func (r *BrandRepository) Delete(ctx context.Context, id string) error {
	_, err := GetExecutor(ctx, r.db).
		Exec(ctx, "DELETE FROM brands WHERE id = $1", id)
	return err
}

func (r *BrandRepository) ExistByName(name string) bool {
	var count int
	err := r.db.
		QueryRow(context.Background(), "SELECT count(*) FROM brands WHERE LOWER(name) = $1", strings.ToLower(name)).
		Scan(&count)
	if err != nil {
		log.Println(err)
		return false
	}
	return count > 0
}

func (r *BrandRepository) Find(params map[string]any) ([]domain.Brand, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM brands;")
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.Brand])
}
