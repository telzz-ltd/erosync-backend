package store

import "github.com/jmoiron/sqlx"

type Store struct {
	User *users
	Otp  *otps
}

func New(db *sqlx.DB) *Store {
	return &Store{
		User: &users{db},
		Otp:  &otps{db},
	}
}
