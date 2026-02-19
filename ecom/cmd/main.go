package main

import (
	"log"

	"github.com/binit2-1/golang-dojo/rest-api/cmd/api"
	"github.com/binit2-1/golang-dojo/rest-api/config"
	psqlDb "github.com/binit2-1/golang-dojo/rest-api/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func main() {
	db, err := psqlDb.NewPSQLStorage(pgx.ConnConfig{
		Config: pgconn.Config{
			Host:     config.Envs.DBAdress,
			User:     config.Envs.DBUser,
			Password: config.Envs.DBPassword,
			Database: config.Envs.DBName,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// Alternative: Using connection string
//
// connStr := "postgres://root:password@localhost:5433/ecom?sslmode=disable"
// db, err := psqlDb.NewPSQLStorage(connStr)
