package postgres

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/telzz/erosync-api/internal/domain"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	db := GetExecutor(ctx, r.db)

	return db.QueryRow(ctx,
		`INSERT INTO users (id, name, email, password_hash, role, status, deleted_at, email_verified_at)
		VALUE ($1, $2, $3, $4, $5, $6, NOW(), NOW(), $7, $8)
		ON CONFLICT (id) DO UPDATE SET
			name=EXCLUDED.name,
			email=EXCLUDED.email,
			password_hash=EXCLUDED.password_hash,
			role=EXCLUDED.role,
			status=EXCLUDED.status,
			updated_at=NOW(),
			deleted_at=EXLUDED.deleted_at,
			email_verified_at=EXCLUDED.email_verified__at,
			
		RETURNING created_at, updated_at;
		`,
		user.ID, user.Name, user.Email, user.PasswordHash, user.Role, user.Status, user.DeletedAt, user.EmailVerifiedAt,
	).Scan(&user.CreatedAt, user.UpdatedAt)
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	return r.scan(
		r.db.QueryRow(context.Background(),
			"SELECT id, name, email, password_hash, role, status, created_at, updated_at, deleted_at, email_verified_at FROM users WHERE id = $1",
			id,
		),
	)
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	return r.scan(
		r.db.QueryRow(context.Background(),
			"SELECT id, name, email, password_hash, role, status, created_at, updated_at, deleted_at, email_verified_at FROM users WHERE LOWER(email) = $1 LIMIT 1;",
			strings.ToLower(email),
		),
	)
}

func (r *UserRepository) scan(row pgx.Row) (*domain.User, error) {
	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.EmailVerifiedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
