package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	validator := &Validator{
		validate: validator.New(),
	}

	err := validator.validate.RegisterValidation("passwd", validatePassword)
	if err != nil {
		panic(err)
	}

	return validator
}

func (v *Validator) ValidateStruct(s interface{}) error {
	return v.validate.Struct(s)
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	uppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	lowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	digit := regexp.MustCompile(`[0-9]`).MatchString(password)
	special := regexp.MustCompile(`[!@#~$%^&*()_+|<>?:{}]`).MatchString(password)

	return uppercase && lowercase && digit && special
}
