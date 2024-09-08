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

	validators.Init()

	conn := database.InitDB(ctx)
	queries := database.New(conn)
	defer conn.Close(ctx)

	router := gin.Default()

	router.Use(
		func(c *gin.Context) {
			c.Set("queries", queries)
			c.Next()
		},
	)

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
