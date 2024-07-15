package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Init() {
	validate = validator.New()
	validate.RegisterValidation("passwd", validatePassword)
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	uppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	lowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	digit := regexp.MustCompile(`[0-9]`).MatchString(password)
	special := regexp.MustCompile(`[!@#~$%^&*()_+|<>?:{}]`).MatchString(password)

	return uppercase && lowercase && digit && special
}
