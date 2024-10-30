package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() *validator.Validate {
	validate = validator.New(validator.WithRequiredStructEnabled())

	return validate
}

func GetValidator() *validator.Validate {
	return validate
}

func MangleValidateResult(err error) map[string]string {
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		errMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {

			// fmt.Println("Namespace: ", err.Namespace())
			// fmt.Println("Tag: ", err.Tag())
			// fmt.Println("Kind: ", err.Kind())
			// fmt.Println("Value: ", err.Value())
			// fmt.Println()

			errMap[err.Namespace()] = fmt.Sprintf("failed on the %s tag, received %v type of %s", err.Tag(), err.Value(), err.Kind().String())
		}

		return errMap
	} else {
		return nil
	}
}
