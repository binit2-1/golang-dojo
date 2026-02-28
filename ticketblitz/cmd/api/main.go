package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"ticketblitz/internal/repository/cache"
	"ticketblitz/internal/repository/pg"
	"ticketblitz/internal/server"

	"github.com/redis/go-redis/v9"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)


func main(){

	err:=godotenv.Load()
	if err!=nil{
		log.Println("Warning: No .env file found, relying on system environment variables")
	}

	dbConnectionURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")


	//POSTGRES CONNECTIONS
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

	//REDIS CONNECTIONS
	redisClient:= redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil{
		log.Fatalf("Error pinging Redis: %v", err)
	}
	fmt.Println("Succesfully connected to Redis")



	pgRepo := pg.NewPotgresEventRepo(db)
	cachedRepo := cache.NewCachedEventRepo(pgRepo, redisClient)


	eventHandler := &server.EventHandler{
		Repo: cachedRepo,
	}

	//ROUTER
	mux := http.NewServeMux()


	//ROUTES
	mux.HandleFunc("POST /v1/events", eventHandler.CreateEvent)
	fmt.Printf("server running on port%s\n", port)


	err = http.ListenAndServe(port, mux)
	if err!=nil{
		log.Fatalf("Failed to start server: %v", err)
	}
}