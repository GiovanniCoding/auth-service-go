package routes

import (
	"github.com/GiovanniCoding/amazon-analysis/auth/app/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.POST("/register", handlers.Register)
}
