package routes

import (
	"github.com/GiovanniCoding/amazon-analysis/auth/app/handlers"
	"github.com/GiovanniCoding/amazon-analysis/auth/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {
	docs.SwaggerInfo.Version = "0.1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Title = "Microservice Auth"
	docs.SwaggerInfo.Description = "Microservice Auth"

	api := router.Group("/api/v1")
	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
