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

	isUserInDB, err := database.Query.UserEmailExist(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check if user exists"})

		return
	}

	if isUserInDB {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})

		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})

		return
	}

	user, err := database.Query.CreateUser(ctx, database.CreateUserParams{
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
