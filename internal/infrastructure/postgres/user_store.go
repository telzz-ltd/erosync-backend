package postgres

import (
	"context"
	"erosync/internal/model"
	"errors"

	"github.com/jackc/pgx/v5"
)

type UserStore struct {
	db DB
}

func (r *UserStore) Save(ctx context.Context, u *model.User) error {
	db := GetDBExecutor(ctx, r.db)
	if u == nil {
		return errors.New("nil user passed")
	}

	_, err := db.Exec(ctx, `
		INSERT INTO users (id, name, email, password_hash, status, role, created_at, updated_at, email_verified_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE SET
			name=EXCLUDED.name,
			email=EXCLUDED.email,
			password_hash=EXCLUDED.password_hash,
			status=EXCLUDED.status,
			role=EXCLUDED.role,
			updated_at=EXCLUDED.updated_at,
			email_verified_at=EXCLUDED.email_verified_at;
	`,
		u.ID,
		u.Name,
		u.Email,
		u.PasswordHash,
		u.Status,
		u.Role,
		u.CreatedAt,
		u.UpdatedAt,
		u.EmailVerifiedAt,
	)

	return err
}

func (r *UserStore) ExistByEmail(email string) (bool, error) {
	var count int
	err := r.db.QueryRow(context.Background(), "SELECT count(*) FROM users WHERE email = $1;", email).Scan(&count)
	return count > 0, err
}

func (r *UserStore) FindByEmail(email string) (*model.User, error) {
	row := r.db.QueryRow(context.Background(), `
		SELECT id, name, email, password_hash, status, role, created_at, updated_at, email_verified_at
		FROM users WHERE email = $1;
	`, email)
	return r.scan(row)
}

func (r *UserStore) FindByID(id string) (*model.User, error) {
	row := r.db.QueryRow(context.Background(), `
		SELECT id, name, email, password_hash, status, role, created_at, updated_at, email_verified_at
		FROM users WHERE email = $1;
	`, id)
	return r.scan(row)
}

func (r *UserStore) scan(row pgx.Row) (*model.User, error) {
	var u model.User
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.PasswordHash,
		&u.Status,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.EmailVerifiedAt,
	)

	return &u, err
}
