package main

import (
	"context"
	"os"

	"github.com/GiovanniCoding/amazon-analysis/auth/app/database"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/middlewares"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/routes"
	"github.com/GiovanniCoding/amazon-analysis/auth/app/validators"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	validators.Init()

	database.InitDB(ctx)
	defer database.Conn.Close(ctx)

	r := gin.Default()

	r.Use(middlewares.LogMiddleware())

	routes.SetupRoutes(r)

	log.Fatal().Err(r.Run())
}
