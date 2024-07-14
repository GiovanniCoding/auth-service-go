package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Procesa la solicitud
		c.Next()

		stop := time.Now()

		log.Info().
			Str("method", c.Request.Method).
			Str("url", c.Request.RequestURI).
			Int("status", c.Writer.Status()).
			Dur("latency", stop.Sub(start)).
			Msg("request")
	}
}
