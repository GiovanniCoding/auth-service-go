package validators_test

import (
	"testing"

	"github.com/GiovanniCoding/auth-microservice/app/validators"
)

type TestStruct struct {
	Password string `validate:"passwd"`
}

func TestValidatePassword(t *testing.T) {
	validator := validators.NewValidator()

	tests := []struct {
		password string
		valid    bool
	}{
		{"Password123!", true},
		{"password123!", false},
		{"PASSWORD123!", false},
		{"Password!!!", false},
		{"Password123", false},
		{"Pass123!", true},
		{"Pass!@#", false},
		{"123456!@", false},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			ts := TestStruct{Password: test.password}

			err := validator.ValidateStruct(ts)

			if (err == nil) != test.valid {
				t.Errorf("expected valid=%v, got valid=%v for password %s", test.valid, err == nil, test.password)
			}
		})
	}
}
