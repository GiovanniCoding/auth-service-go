package routes

import (
	"github.com/GiovanniCoding/auth-microservice/app/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	api.POST("/login", handlers.Login)
	api.POST("/register", handlers.Register)

	api.POST("/validate-token", handlers.ValidateToken)
}
