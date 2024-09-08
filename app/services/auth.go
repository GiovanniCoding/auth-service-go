package services

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GiovanniCoding/auth-microservice/app/database"
	"github.com/GiovanniCoding/auth-microservice/app/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterProcess(request schemas.RegisterRequest, ctx *gin.Context) (schemas.RegisterResponse, error) {
	var response schemas.RegisterResponse

	queries, ok := ctx.MustGet("queries").(*database.Queries)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})

		return response, errors.New("database connection not found")
	}

	isUserInDB, err := queries.UserEmailExist(ctx, request.Email)
	if err != nil {
		return response, errors.New("failed to check if user exists")
	}

	if isUserInDB {
		return response, errors.New("user already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, errors.New("failed to hash password")
	}

	user, err := queries.CreateUser(ctx, database.CreateUserParams{
		Email:        request.Email,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		return response, errors.New("failed to create user")
	}

	response.ID = user.ID
	response.Email = user.Email

	return response, nil
}

func LoginProcess(request schemas.LoginRequest, ctx *gin.Context) (schemas.LoginResponse, error) {
	var response schemas.LoginResponse

	queries, ok := ctx.MustGet("queries").(*database.Queries)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})

		return response, errors.New("database connection not found")
	}

	user, err := queries.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return response, errors.New("user or password incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return response, errors.New("user or password incorrect")
	}

	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := schemas.LoginClaim{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Printf("Error signing token: %v\n", err)

		return response, errors.New("failed to sign token")
	}

	response.Token = signedToken

	return response, nil
}

func ValidateTokenProcess(request schemas.ValidateTokenRequest, ctx *gin.Context) (bool, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(request.Token, &schemas.LoginClaim{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return false, errors.New("invalid token")
	}

	_, ok := token.Claims.(*schemas.LoginClaim)
	if !ok || !token.Valid {
		return false, errors.New("invalid token")
	}

	return true, nil
}
