package schemas

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type RegisterUserRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,passwd"`
}

type RegisterUserResponse struct {
	ID    pgtype.UUID `json:"id"`
	Email string      `json:"email"`
}
