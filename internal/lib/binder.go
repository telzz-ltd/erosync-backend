package lib

import "github.com/go-playground/validator/v10"

type Binder struct {
	validate *validator.Validate
}

func NewBinder() *Binder {
	return &Binder{
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
