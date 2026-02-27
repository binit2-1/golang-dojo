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
		fmt.Println("Couldn't Load env")
	}

	dbConnectionURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("pgx", dbConnectionURL)
	// sql.Open() only validates arguments, so this catches basic format errors
	if err != nil{
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	// This is where you catch actual connection errors (e.g., wrong credentials, network issues)
	if err != nil{
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")
}