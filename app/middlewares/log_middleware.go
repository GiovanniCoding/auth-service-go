package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func LogMiddleware(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		stop := time.Now()

		log.Info().
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("status", c.Response().StatusCode()).
			Dur("latency", stop.Sub(start)).
			Msg("request")

		return err
	})
}
