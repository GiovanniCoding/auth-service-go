package schemas

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type RegisterRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,passwd"`
}

type RegisterResponse struct {
	ID    pgtype.UUID `json:"id"`
	Email string      `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,passwd"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginClaim struct {
	UserID pgtype.UUID `json:"userId"`
	jwt.RegisteredClaims
}

type ValidateTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

type ValidateTokenResponse struct {
	Valid bool `json:"valid"`
}
