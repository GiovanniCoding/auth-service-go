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

// @BasePath /api/v1

// register godoc
// @Summary Create New User
// @Schemes
// @Description Create New User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body schemas.RegisterUserRequest true "New user info"
// @Success 201 {object} schemas.RegisterUserResponse "New user created"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /register [post]
func Register(ctx *gin.Context) {
	var req schemas.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "invalid request"})

		return
	}

	if err := validators.ValidateStruct(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "Validation error: " + err.Error()})

		return
	}

	isUserInDB, err := database.Query.UserEmailExist(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: "failed to check if user exists"})

		return
	}

	if isUserInDB {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})

		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: "failed to hash password"})

		return
	}

	user, err := database.Query.CreateUser(ctx, database.CreateUserParams{
		Email:        req.Email,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: "failed to create user"})

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
