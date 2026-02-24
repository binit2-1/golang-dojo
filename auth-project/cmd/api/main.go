package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/binit2-1/golang-dojo/auth-project/internal/database"
	"github.com/binit2-1/golang-dojo/auth-project/internal/middleware"
	"github.com/binit2-1/golang-dojo/auth-project/internal/server"
)

func main(){

	rdb := database.NewRedisClient() 
	defer rdb.Close()

	srv := server.NewServer(rdb)

	limiter := middleware.NewRateLimiter(rdb)
	loginHandler := limiter.Limit(http.HandlerFunc(srv.LoginHandler))

	
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"messages":"auth-project up"}`))
	})

	mux.Handle("POST /v1/login", loginHandler)
	mux.Handle("GET /v1/dashboard", srv.RequireAuth(http.HandlerFunc(srv.DashBoardHandler)))

	wrappedMux := middleware.LogRequest(mux)
	port := ":8080"
	fmt.Printf("Starting Server on port %s\n", port)


	err := http.ListenAndServe(port, wrappedMux)
	if err != nil{
		log.Fatalf("Failed to start server: %v", err)
	}

	
}
