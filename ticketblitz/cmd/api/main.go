package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"ticketblitz/internal/repository/cache"
	"ticketblitz/internal/repository/pg"
	"ticketblitz/internal/server"

	"github.com/redis/go-redis/v9"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	// =========================================================================
	// 1. CONFIGURATION & ENVIRONMENT
	// =========================================================================
	if err := godotenv.Load(); err != nil {
		log.Println("WARN: No .env file found. Relying on system environment variables.")
	}

	dbConnectionURL := os.Getenv("DATABASE_URL")
	if dbConnectionURL == "" {
		log.Fatal("FATAL: DATABASE_URL environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // Strict fallback to prevent silently binding to port 80
	}

	// =========================================================================
	// 2. POSTGRESQL CONNECTION (The Source of Truth)
	// =========================================================================
	db, err := sql.Open("pgx", dbConnectionURL)
	if err != nil {
		log.Fatalf("FATAL: Failed to parse Postgres configuration: %v", err)
	}
	defer db.Close() // Guarantees connection pool is closed if main() exits

	// Create a strict 5-second timeout context for the Ping. 
	// If the DB doesn't answer in 5 seconds, we crash the app immediately.
	pgCtx, pgCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pgCancel()

	if err := db.PingContext(pgCtx); err != nil {
		log.Fatalf("FATAL: Postgres is unreachable: %v", err)
	}
	fmt.Println("✅ Successfully connected to Postgres!")

	// =========================================================================
	// 3. REDIS CONNECTION (The Speed Layer)
	// =========================================================================
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // In production, this comes from an env variable
	})
	defer redisClient.Close()

	// Strict 5-second timeout for Redis
	redisCtx, redisCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer redisCancel()

	if err := redisClient.Ping(redisCtx).Err(); err != nil {
		log.Fatalf("FATAL: Redis is unreachable: %v", err)
	}
	fmt.Println("✅ Successfully connected to Redis!")

	// =========================================================================
	// 4. DEPENDENCY INJECTION (Wiring the layers together)
	// =========================================================================
	// Instantiate the core Postgres repository
	pgRepo := pg.NewPostgresEventRepo(db)
	
	// Wrap the core repository with the Redis Cache Decorator
	cachedRepo := cache.NewCachedEventRepo(pgRepo, redisClient)

	// Inject the wrapped repository into the HTTP Handler
	eventHandler := &server.EventHandler{
		Repo: cachedRepo,
	}

	// =========================================================================
	// 5. ROUTING & HTTP SERVER
	// =========================================================================
	mux := http.NewServeMux()

	// Register API endpoints
	mux.HandleFunc("POST /v1/events", eventHandler.CreateEvent)
	mux.HandleFunc("GET /v1/events/{id}", eventHandler.GetEventByID)

	// Start the server
	fmt.Printf("🚀 Server running on port %s\n", port)
	
	// ListenAndServe blocks forever unless it crashes
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("FATAL: Server crashed: %v", err)
	}
}