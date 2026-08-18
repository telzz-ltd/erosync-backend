package api

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func RegisterValidator() {
	Validate = validator.New()

	Validate.RegisterValidation("name", func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(string)
		exp, err := regexp.Compile(`^[a-zA-Z]{3,}(?:\s[a-zA-Z]{3,}){1,2}$`)
		if err != nil {
			return false
		}

		return exp.MatchString(value)

	})
}
