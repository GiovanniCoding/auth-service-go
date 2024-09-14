package schemas_test

import (
	"testing"

	"github.com/GiovanniCoding/auth-microservice/app/schemas"
	"github.com/GiovanniCoding/auth-microservice/app/validators"
)

func TestRegisterUserRequest(t *testing.T) {
	validator := validators.NewValidator()

	tests := []struct {
		email    string
		password string
		valid    bool
	}{
		{"test@example.com", "Password123!", true},
		{"invalid-email", "Password123!", false},
		{"test@example.com", "short", false},
		{"", "Password123!", false},
		{"test@example.com", "", false},
		{"test@example.com", "password123!", false},
		{"test@example.com", "PASSWORD123!", false},
		{"test@example.com", "Password!!!", false},
		{"test@example.com", "Pass123!", true},
	}

	for _, test := range tests {
		t.Run(
			test.email+"_"+test.password,
			func(t *testing.T) {
				req := schemas.RegisterRequest{
					Email:    test.email,
					Password: test.password,
				}

				err := validator.ValidateStruct(&req)

				if (err == nil) != test.valid {
					t.Errorf(
						"expected valid=%v, got valid=%v for email %s and password %s", test.valid, err == nil, test.email, test.password,
					)
				}
			},
		)
	}
}
