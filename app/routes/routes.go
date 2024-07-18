package routes

import (
	"github.com/GiovanniCoding/amazon-analysis/auth/app/handlers"
	"github.com/GiovanniCoding/amazon-analysis/auth/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"

	api := router.Group("/api/v1")
	api.POST("/register", handlers.Register)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
