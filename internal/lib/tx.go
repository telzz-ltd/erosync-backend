package lib

import "context"

type Tx interface {
	Execute(context.Context, func(context.Context) error) error
}
