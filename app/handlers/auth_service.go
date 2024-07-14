package handlers

import (
	"context"

	"github.com/GiovanniCoding/amazon-analysis/auth/app/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	data := map[string]string{
		"email":    "email",
		"password": "password",
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	user, err := queries.CreateUser(ctx, database.CreateUserParams{
		Email:        data["email"],
		PasswordHash: string(password),
	})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
