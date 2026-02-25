package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/binit2-1/golang-dojo/microservices-jwt/internal/middlewares"
	"github.com/joho/godotenv"
)



func main(){


	mux := http.NewServeMux()
	wrappedMux := middlewares.LogRequest(mux)


	http.HandleFunc("GET /v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"messages:"jwt project up"}`))
	})


	godotenv.Load()

	port := ":" + os.Getenv("PORT")
	fmt.Printf("Starting Server on port%s\n", port)


	err := http.ListenAndServe(port, wrappedMux)
	if err != nil{
		log.Fatalf("Failed to start server: %v", err)
	}
}

