package app

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func RegisterValidator() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
