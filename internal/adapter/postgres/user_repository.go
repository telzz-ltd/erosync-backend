package postgres

import (
	"context"

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
		VALUE (@id, @name, @email, @password_hash, @role, @status, NOW(), NOW(), @deleted_at, @email_verified_at)
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
		pgx.StructArgs(&user),
	).Scan(&user.CreatedAt, user.UpdatedAt)
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRow(context.Background(),
		"SELECT * FROM users WHERE id = $1;",
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.EmailVerifiedAt)

	return &user, err
}

func (r *UserRepository) scan(row pgx.Row) (*domain.User, error) {
	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.EmailVerifiedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
