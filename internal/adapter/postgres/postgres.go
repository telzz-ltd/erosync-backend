package postgres

import (
	"context"
	"erosync/internal/port"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var executorKey = "pg-tx"

type DBTX interface {
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
}

func GetExecutor(ctx context.Context, fallback DBTX) DBTX {
	executor, ok := ctx.Value(executorKey).(DBTX)
	if !ok {
		return fallback
	}

	return executor
}

func New(dbUrl string) *port.Store {
	db, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalln("uanble to connect to db", err)
	}

	return &port.Store{
		User: NewUserRepository(db),
		Otp:  NewOTPRepository(db),
		Tx:   NewTx(db),
	}
}
