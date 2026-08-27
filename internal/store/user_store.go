package store

import (
	"erosync/internal/model"

	"github.com/jmoiron/sqlx"
)

type users struct {
	db *sqlx.DB
}

func (u *users) Save(user *model.User) error {
	sql := `
		INSERT into users 
			(id, name, email, password_hash, status, role, created_at, updated_at, email_verified_at)
		VALUES 
			(:id, :name, :email, :password_hash, :status, :role, :created_at, :updated_at, :email_verified_at)
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
	`
	_, err := u.db.NamedExec(sql, user)
	return err
}

func (u *users) FindByID(id string) (*model.User, error) {
	var user model.User
	err := u.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *users) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
