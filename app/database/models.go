// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID           pgtype.UUID
	Email        string
	PasswordHash string
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
	DeletedAt    pgtype.Timestamptz
}
