package validators

import (
	"testing"
)

type TestStruct struct {
	Password string `validate:"passwd"`
}

func TestValidatePassword(t *testing.T) {
	// Inicializar el validador
	Init()

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
			// Crear una estructura de prueba con la contraseña
			ts := TestStruct{Password: test.password}

			// Validar la estructura
			err := ValidateStruct(ts)

			if (err == nil) != test.valid {
				t.Errorf("expected valid=%v, got valid=%v for password %s", test.valid, err == nil, test.password)
			}
		})
	}
}
