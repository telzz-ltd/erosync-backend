package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

type TxManager struct {
	db *pgxpool.Pool
}

func (m *TxManager) Transaction(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	tx, err := m.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	ctx = context.WithValue(ctx, "tx", tx)
	err = fn(ctx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func GetDBExecutor(ctx context.Context, fallback DB) DB {
	tx, ok := ctx.Value("tx").(DB)
	if !ok {
		return fallback
	}

	return tx
}
