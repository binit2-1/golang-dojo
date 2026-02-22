package main

import (
	// "database/sql/driver"
	"fmt"
	"log"

	"github.com/binit2-1/golang-dojo/ecom/config"
	psqlDb "github.com/binit2-1/golang-dojo/ecom/db"
	// "github.com/golang-migrate/migrate/v4"
)

func main(){
	db, err := psqlDb.NewPSQLStorage(config.Envs.DBConnStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hi", db)
}
