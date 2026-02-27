package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/jackc/pgx/v5/stdlib"
)


func main(){

	err:=godotenv.Load()
	if err!=nil{
		log.Println("Warning: No .env file found, relying on system environment variables")
	}

	dbConnectionURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("pgx", dbConnectionURL)
	// sql.Open() only validates arguments, so this catches basic format errors
	if err != nil{
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	// catch actual connection errors (e.g., wrong credentials, network issues)
	if err != nil{
		log.Fatalf("Error pinging database: %v", err)
		
	}
	fmt.Println("Successfully connected to the database!")
}