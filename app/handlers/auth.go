package handlers

import (
	"net/http"

	"github.com/GiovanniCoding/amazon-analysis/auth/app/schemas"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/services"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/validators"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var request schemas.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "invalid request"})

		return
	}

	if err := validators.ValidateStruct(&request); err != nil {
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
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "invalid request"})

		return
	}

	if err := validators.ValidateStruct(&request); err != nil {
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
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "invalid request"})

		return
	}

	if err := validators.ValidateStruct(&request); err != nil {
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
