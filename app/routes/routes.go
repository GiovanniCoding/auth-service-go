package routes

import (
	"github.com/GiovanniCoding/amazon-analysis/auth/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/register", handlers.Register)
}
