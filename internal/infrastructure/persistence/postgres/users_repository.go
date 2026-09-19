package postgres

import (
	"context"
	"erosync/internal/users"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Save(ctx context.Context, user users.User) error {
	db := GetExecutor(ctx, r.db)
	_, err := db.Exec(ctx,
		`
			INSERT into users
				(id, name, email, password_hash, status, role, created_at, updated_at, email_verified_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			ON CONFLICT (id) DO UPDATE SET
				name=EXCLUDED.name,
				email=EXCLUDED.email,
				password_hash=EXCLUDED.password_hash,
				status=EXCLUDED.status,
				role=EXCLUDED.role,
				created_at=EXCLUDED.created_at,
				updated_at=EXCLUDED.updated_at,
				email_verified_at=EXCLUDED.email_verified_at
			;
		`,
		user.ID,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Status,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
		user.EmailVerifiedAt,
	)

	return err
}

func (r *UserRepository) FindByID(id string) (users.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM users WHERE id = $1;", id)
	if err != nil {
		return users.User{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[users.User])
}

func (r *UserRepository) FindByEmail(email string) (users.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM users WHERE email = $1;", email)
	if err != nil {
		return users.User{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[users.User])
}
