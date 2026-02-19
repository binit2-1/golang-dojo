package db

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

// NewPSQLStorage creates a new PostgreSQL connection using a connection string.
// The connStr is parsed via pgx.ParseConfig (required by pgx internally),
// then registered with stdlib for use with database/sql.
func NewPSQLStorage(connStr string) (*sql.DB, error) {
	cfg, err := pgx.ParseConfig(connStr)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("pgx", stdlib.RegisterConnConfig(cfg))
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
