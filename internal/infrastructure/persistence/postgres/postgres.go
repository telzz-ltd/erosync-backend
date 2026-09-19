package postgres

import (
	"context"

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

type Tx struct {
	db *pgxpool.Pool
}

func (tx Tx) Execute(ctx context.Context, cb func(ctx context.Context) error) error {
	_tx, err := tx.db.Begin(ctx)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, executorKey, _tx)

	if err := cb(ctx); err != nil {
		_tx.Rollback(ctx)
		return err
	}

	return _tx.Commit(ctx)
}
