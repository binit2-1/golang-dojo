package main

import (
	"database/sql"
	"log"

	"github.com/binit2-1/golang-dojo/rest-api/cmd/api"
	"github.com/binit2-1/golang-dojo/rest-api/config"
	psqlDb "github.com/binit2-1/golang-dojo/rest-api/db"
)

func main() {
	db, err := psqlDb.NewPSQLStorage(config.Envs.DBConnStr)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully Connected")
}
