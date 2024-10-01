package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

var validate *validator.Validate

func InitValidator() *validator.Validate {
	validate = validator.New(validator.WithRequiredStructEnabled())

	return validate
}

func GetValidator() *validator.Validate {
	return validate
}

func RegisterCustomValidations(v *validator.Validate) {
	v.RegisterValidation("specialDiag", models.ValidateSpecialReferralDiagnosisCode)
}
