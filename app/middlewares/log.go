package middlewares

import (
	"bytes"
	"io"
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

		var requestBody []byte
		if ctx.Request.Body != nil {
			requestBody, _ = io.ReadAll(ctx.Request.Body)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		Logger.Info().
			Str("url", ctx.Request.RequestURI).
			Str("method", ctx.Request.Method).
			Str("client_ip", ctx.ClientIP()).
			Bytes("request_body", requestBody).
			Msg("request")

		start := time.Now()

		ctx.Next()

		stop := time.Now()

		Logger.Info().
			Int("status", ctx.Writer.Status()).
			Dur("latency", stop.Sub(start)).
			Int("response_size", ctx.Writer.Size()).
			Msg("response")
	}
}
