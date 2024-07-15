package handlers

import (
	"net/http"

	"github.com/GiovanniCoding/amazon-analysis/auth/app/database"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/middlewares"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/schemas"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/validators"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var req schemas.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := validators.ValidateStruct(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation error: " + err.Error()})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user, err := database.Q.CreateUser(ctx, database.CreateUserParams{
		Email:        req.Email,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	middlewares.Logger.Info().
		Str("email", user.Email).
		Msg("User created successfully")

	ctx.JSON(http.StatusCreated, schemas.RegisterUserResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}
