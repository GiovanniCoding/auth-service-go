package routes

import (
	"github.com/GiovanniCoding/auth-microservice/app/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	api.GET("/health", handlers.HealthCheck)

	api.POST("/signup", handlers.Signup)
	api.POST("/login", handlers.Login)

	api.POST("/validate-token", handlers.ValidateToken)
}
