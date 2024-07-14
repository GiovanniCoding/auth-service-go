package main

import (
	"flag"
	"os"

	"github.com/GiovanniCoding/amazon-analysis/auth/app/middlewares"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Parse command-line flags
	flag.Parse()

	// Create Gin app
	r := gin.Default()

	// Middleware
	r.Use(middlewares.LogMiddleware())

	// Setup routes
	routes.SetupRoutes(r)

	// Listen on port 3000
	log.Fatal().Err(r.Run())
}
