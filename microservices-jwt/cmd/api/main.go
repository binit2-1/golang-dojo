package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/binit2-1/golang-dojo/microservices-jwt/internal/middlewares"
)



func main(){


	mux := http.NewServeMux()
	wrappedMux := middlewares.LogRequest(mux)


	http.HandleFunc("GET /v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"messages:"jwt project up"}`))
	})


	port := ":8080"
	fmt.Printf("Starting Server on port%s\n", port)


	err := http.ListenAndServe(port, wrappedMux)
	if err != nil{
		log.Fatalf("Failed to start server: %v", err)
	}
}

