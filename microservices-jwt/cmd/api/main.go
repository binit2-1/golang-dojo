package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/binit2-1/golang-dojo/microservices-jwt/internal/middlewares"
	"github.com/binit2-1/golang-dojo/microservices-jwt/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No env was found")
	}

	mux := http.NewServeMux()
	wrappedMux := middlewares.LogRequest(mux)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	port := ":" + os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	srv := server.NewServer(secret)

	mux.HandleFunc("GET /v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"messages:"jwt project up"}`))
	})
	mux.HandleFunc("POST /v1/login", srv.LoginHandler)
	mux.Handle("GET /v1/dashboard", srv.RequireJWT(http.HandlerFunc(srv.DashboardHandler)))

	fmt.Printf("Starting Server on port%s\n", port)

	err = http.ListenAndServe(port, wrappedMux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
