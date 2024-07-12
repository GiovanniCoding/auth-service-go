package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", Home)
}

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
