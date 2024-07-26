package database

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/jackc/pgx/v5"
)

func InitDB(ctx context.Context) *pgx.Conn {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Panic("DATABASE_URL is required")
	}

	conn, err := pgx.Connect(ctx, dsn)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	runMigrations()
	log.Println("Connected to the database")

	return conn
}

func runMigrations() {
	applyCmd := exec.Command("atlas", "migrate", "apply", "--url", os.Getenv("DATABASE_URL"))

	if output, err := applyCmd.CombinedOutput(); err != nil {
		log.Fatalf("Error aplicando migraciones: %s\n%s", err, output)
	}

	log.Println("Migraciones aplicadas correctamente.")
}
