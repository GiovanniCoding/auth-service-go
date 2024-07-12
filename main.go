package main

import (
	"flag"
	"os"

	"github.com/GiovanniCoding/amazon-analysis/auth/app/middlewares"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var port = flag.String("port", ":3000", "Port to listen on")

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Parse command-line flags
	flag.Parse()

	// Create fiber app
	app := fiber.New(fiber.Config{})

	// Middleware
	middlewares.LogMiddleware(app)

	// Setup routes
	routes.SetupRoutes(app)

	// // Create a /api/v1 endpoint
	// v1 := app.Group("/api/v1")

	// // Bind handlers
	// v1.Get("/users", handlers.UserList)
	// v1.Post("/users", handlers.UserCreate)

	// Listen on port 3000
	log.Fatal().Err(app.Listen(*port))
}
