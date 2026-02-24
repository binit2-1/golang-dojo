package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/binit2-1/golang-dojo/middleware-chaining/internal/middleware"
)


func main(){
	//new router
	mux := http.NewServeMux()

	//health route
	mux.HandleFunc("GET /v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Health OK"}`))
	} )
	
	//Rate limiter
	limiter := middleware.NewRateLimiter()

	//wrapped the mux with the middleware
	wrappedMux := middleware.LogRequest(limiter.Limit(mux))

	//port
	port := ":8080"
	fmt.Printf("Starting Server on port %s\n", port)
	
	err := http.ListenAndServe(port, wrappedMux)
	if err != nil{
		log.Fatalf("Failed to start server: %v", err)
	}

}