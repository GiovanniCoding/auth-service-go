package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		c.Set("traceID", traceID)
		Logger = log.With().Str("trace_id", traceID).Logger()

		Logger.Info().
			Str("url", c.Request.RequestURI).
			Str("method", c.Request.Method).
			Msg("request")

		start := time.Now()

		c.Next()

		stop := time.Now()

		Logger.Info().
			Int("status", c.Writer.Status()).
			Dur("latency", stop.Sub(start)).
			Msg("response")
	}
}
