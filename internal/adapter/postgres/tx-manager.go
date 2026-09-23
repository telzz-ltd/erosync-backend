package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Tx struct {
	db *pgxpool.Pool
}

func NewTx(db *pgxpool.Pool) *Tx {
	return &Tx{db}
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
