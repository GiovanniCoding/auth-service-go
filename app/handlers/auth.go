package handlers

import (
	"net/http"

	"github.com/GiovanniCoding/amazon-analysis/auth/app/middlewares"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/schemas"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/services"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/validators"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// register godoc
// @Summary Create New User
// @Schemes
// @Description Create New User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body schemas.RegisterRequest true "New user info"
// @Success 201 {object} schemas.RegisterResponse "New user created"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /register [post]
func Register(ctx *gin.Context) {
	var request schemas.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		middlewares.Logger.Error().
			Msg("invalid request")
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "invalid request"})

		return
	}

	if err := validators.ValidateStruct(&request); err != nil {
		middlewares.Logger.Error().
			Msg("Validation error: " + err.Error())
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

// login godoc
// @Summary Login User
// @Schemes
// @Description Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body schemas.LoginRequest true "User login info"
// @Success 200 {object} schemas.LoginResponse "User logged in"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /login [post]
func Login(ctx *gin.Context) {
	var request schemas.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		middlewares.Logger.Error().
			Msg("invalid request")
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: "invalid request"})

		return
	}

	if err := validators.ValidateStruct(&request); err != nil {
		middlewares.Logger.Error().
			Msg("Validation error: " + err.Error())
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
