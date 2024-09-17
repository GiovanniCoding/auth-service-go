package services

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GiovanniCoding/auth-microservice/app/database"
	errs "github.com/GiovanniCoding/auth-microservice/app/errors"
	"github.com/GiovanniCoding/auth-microservice/app/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignupProcess(request schemas.SignupRequest, ctx *gin.Context) (schemas.SignupResponse, error) {
	var response schemas.SignupResponse

	queries, ok := ctx.MustGet("queries").(*database.Queries)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})

		return response, errs.ErrDBConnection
	}

	isUserInDB, err := queries.UserEmailExist(ctx, request.Email)
	if err != nil {
		return response, errs.ErrDBConnection
	}

	if isUserInDB {
		return response, errs.ErrUserAlreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, errs.ErrInternalServerErr
	}

	user, err := queries.CreateUser(ctx, database.CreateUserParams{
		Email:        request.Email,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		return response, errs.ErrDBConnection
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

		return response, errs.ErrDBConnection
	}

	user, err := queries.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return response, errs.ErrUserPassInvalid
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return response, errs.ErrUserPassInvalid
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

		return response, errs.ErrInternalServerErr
	}

	response.Token = signedToken

	return response, nil
}

func ValidateTokenProcess(request schemas.ValidateTokenRequest) (bool, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(
		request.Token,
		&schemas.LoginClaim{},
		func(_ *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)
	if err != nil {
		return false, errs.ErrInvalidToken
	}

	_, ok := token.Claims.(*schemas.LoginClaim)
	if !ok || !token.Valid {
		return false, errs.ErrInvalidToken
	}

	return true, nil
}
