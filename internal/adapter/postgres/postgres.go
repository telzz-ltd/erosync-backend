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
	ctx := context.Background()

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		panic(err)
	}

	// config.ConnConfig.Tracer = &tracelog.TraceLog{
	// 	Logger: tracelog.LoggerFunc(func(
	// 		ctx context.Context,
	// 		level tracelog.LogLevel,
	// 		msg string,
	// 		data map[string]any,
	// 	) {
	// 		log.Printf("[%s] %s %v", level, msg, data)
	// 	}),
	// 	LogLevel: tracelog.LogLevelTrace,
	// }

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalln("uanble to connect to db", err)
	}

	return &port.Store{
		Tx: NewTx(db),

		User:          NewUserRepository(db),
		Otp:           NewOTPRepository(db),
		Brand:         NewBrandRepository(db),
		BrandCategory: NewBrandCategoryRepository(db),
	}
}
