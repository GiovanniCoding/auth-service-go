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
	return func(ctx *gin.Context) {
		traceID := uuid.New().String()
		ctx.Set("traceID", traceID)
		Logger = log.With().Str("trace_id", traceID).Logger()

		Logger.Info().
			Str("url", ctx.Request.RequestURI).
			Str("method", ctx.Request.Method).
			Msg("request")

		start := time.Now()

		ctx.Next()

		stop := time.Now()

		Logger.Info().
			Int("status", ctx.Writer.Status()).
			Dur("latency", stop.Sub(start)).
			Msg("response")
	}
}
