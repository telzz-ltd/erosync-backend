package postgres

import (
	"context"
	"erosync/internal/store"
	"erosync/pkg/env"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New() store.Store {
	db, err := pgxpool.New(context.Background(), env.GetOrPanic("DATABASE_URL"))
	if err != nil {
		log.Panicln("Unable to connect to db:", err)
	}

	return store.Store{
		User: &UserStore{db},
	}
}
