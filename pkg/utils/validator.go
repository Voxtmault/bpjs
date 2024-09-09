package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func InitValidator() *validator.Validate {
	validate = validator.New(validator.WithRequiredStructEnabled())

	return validate
}

func GetValidator() *validator.Validate {
	return validate
}
