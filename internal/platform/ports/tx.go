package ports

import "context"

type TxExecutor interface {
	Execute(context.Context, func(context.Context) error) error
}
