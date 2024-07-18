package services

import (
	"errors"

	"github.com/GiovanniCoding/amazon-analysis/auth/app/database"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/middlewares"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterProcess(request schemas.RegisterUserRequest, ctx *gin.Context) (schemas.RegisterUserResponse, error) {
	var response schemas.RegisterUserResponse

	isUserInDB, err := database.Query.UserEmailExist(ctx, request.Email)
	if err != nil {
		middlewares.Logger.Error().
			Msg("failed to check if user exists")
		return response, errors.New("failed to check if user exists")
	}

	if isUserInDB {
		middlewares.Logger.Error().
			Msg("user already exists")
		return response, errors.New("user already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		middlewares.Logger.Error().
			Msg("failed to hash password")
		return response, errors.New("failed to hash password")
	}

	user, err := database.Query.CreateUser(ctx, database.CreateUserParams{
		Email:        request.Email,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		middlewares.Logger.Error().
			Msg("failed to create user")
		return response, errors.New("failed to create user")
	}

	middlewares.Logger.Info().
		Str("email", user.Email).
		Msg("User created successfully")

	response.ID = user.ID
	response.Email = user.Email
	return response, nil
}
