package routes

import (
	"github.com/GiovanniCoding/amazon-analysis/auth/app/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// api := app.Group("/api/v1")
	r.POST("/register", handlers.Register)
}
