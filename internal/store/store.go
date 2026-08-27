package store

import "github.com/jmoiron/sqlx"

type Store struct {
	User *users
}

func New(db *sqlx.DB) *Store {
	return &Store{
		User: &users{db},
	}
}
