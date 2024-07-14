package schemas

import "github.com/jackc/pgx/v5/pgtype"

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	ID    pgtype.UUID `json:"id"`
	Email string      `json:"email"`
}
