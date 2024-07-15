package database

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn
var Query *Queries

func InitDB(ctx context.Context) {
	var err error
	Conn, err = pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	runMigrations()

	Query = New(Conn)

	log.Println("Connected to the database")
}

func runMigrations() {
	envVars := []string{
		"DB_URL=" + os.Getenv("DB_URL"),
		"MIGRATIONS_PATH=" + os.Getenv("MIGRATIONS_PATH"),
		"SCHEMA_PATH=" + os.Getenv("SCHEMA_PATH"),
		"DB_DEV_URL=" + os.Getenv("DB_DEV_URL"),
	}

	applyCmd := exec.Command("atlas", "migrate", "apply", "--url", os.Getenv("DB_URL"))
	applyCmd.Env = append(os.Environ(), envVars...)

	if output, err := applyCmd.CombinedOutput(); err != nil {
		log.Fatalf("Error aplicando migraciones: %s\n%s", err, output)
	}

	log.Println("Migraciones aplicadas correctamente.")
}
