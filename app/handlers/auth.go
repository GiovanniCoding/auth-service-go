package handlers

import (
	"net/http"

	"github.com/GiovanniCoding/auth-microservice/app/schemas"
	"github.com/GiovanniCoding/auth-microservice/app/services"
	"github.com/GiovanniCoding/auth-microservice/app/validators"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var request schemas.RegisterRequest

	validatorInterface, validatorExists := ctx.Get("validator")
	if !validatorExists {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: "validator not found"})

		return
	}

	validator, isValidator := validatorInterface.(*validators.Validator)
	if !isValidator {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: "invalid validator type"})

		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "invalid request"})

		return
	}

	if err := validator.ValidateStruct(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "Validation error: " + err.Error()})

		return
	}

	response, err := services.RegisterProcess(request, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func Login(ctx *gin.Context) {
	var request schemas.LoginRequest

	validatorInterface, validatorExists := ctx.Get("validator")
	if !validatorExists {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: "validator not found"})

		return
	}

	validator, isValidator := validatorInterface.(*validators.Validator)
	if !isValidator {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: "invalid validator type"})

		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "invalid request"})

		return
	}

	if err := validator.ValidateStruct(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "Validation error: " + err.Error()})

		return
	}

	response, err := services.LoginProcess(request, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, response)
}

func ValidateToken(ctx *gin.Context) {
	var request schemas.ValidateTokenRequest

	validatorInterface, validatorExists := ctx.Get("validator")
	if !validatorExists {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: "validator not found"})

		return
	}

	validator, isValidator := validatorInterface.(*validators.Validator)
	if !isValidator {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: "invalid validator type"})

		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "invalid request"})

		return
	}

	if err := validator.ValidateStruct(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "Validation error: " + err.Error()})

		return
	}

	isValid, err := services.ValidateTokenProcess(request, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: err.Error()})

		return
	}

	if !isValid {
		ctx.JSON(http.StatusUnauthorized, schemas.ValidateTokenResponse{Valid: false})

		return
	}

	ctx.JSON(http.StatusOK, schemas.ValidateTokenResponse{Valid: true})
}
