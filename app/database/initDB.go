package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn
var Q *Queries

func InitDB(ctx context.Context) {
	var err error
	Conn, err = pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	Q = New(Conn)
	log.Println("Connected to the database")
}
