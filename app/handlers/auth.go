package handlers

import (
	"context"
	"net/http"

	"github.com/GiovanniCoding/amazon-analysis/auth/app/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	data := map[string]string{
		"email":    "email",
		"password": "password",
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	_, err = queries.CreateUser(ctx, database.CreateUserParams{
		Email:        data["email"],
		PasswordHash: string(password),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusCreated)
}
