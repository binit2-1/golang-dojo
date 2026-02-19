package db

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

func NewPSQLStorage(cfg pgx.ConnConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", stdlib.RegisterConnConfig(&cfg))
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

// Alternative: Using connection string
//
// import (
// 	"database/sql"
// 	"log"
// 	_ "github.com/jackc/pgx/v5/stdlib"
// )
//
// func NewPSQLStorage(connStr string) (*sql.DB, error) {
// 	db, err := sql.Open("pgx", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db, nil
// }
//
// Usage:
// connStr := "postgres://root:password@localhost:5433/ecom?sslmode=disable"
// db, err := db.NewPSQLStorage(connStr)
