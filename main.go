package main

import (
	"context"
	"log"

	"github.com/GiovanniCoding/auth-microservice/app/database"
	"github.com/GiovanniCoding/auth-microservice/app/routes"
	"github.com/GiovanniCoding/auth-microservice/app/validators"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	validator := validators.NewValidator()

	conn := database.InitDB(ctx)
	queries := database.New(conn)
	defer conn.Close(ctx)

	router := gin.Default()

	router.Use(
		func(c *gin.Context) {
			c.Set("queries", queries)
			c.Set("validator", validator)
			c.Next()
		},
	)

	routes.SetupRoutes(router)

	if err := router.Run(":30001"); err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
