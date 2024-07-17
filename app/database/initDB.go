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
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL no está configurado en las variables de entorno")
	}

	Conn, err := pgx.Connect(ctx, dbURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	runMigrations()

	Query = New(Conn)

	log.Println("Connected to the database")
}

func runMigrations() {
	applyCmd := exec.Command("atlas", "migrate", "apply", "--url", os.Getenv("DATABASE_URL"))

	if output, err := applyCmd.CombinedOutput(); err != nil {
		log.Fatalf("Error aplicando migraciones: %s\n%s", err, output)
	}

	log.Println("Migraciones aplicadas correctamente.")
}
