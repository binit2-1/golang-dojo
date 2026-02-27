package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"realtime-canvas/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()

	port := os.Getenv("PORT")

	mux.HandleFunc("/ws", server.ServeWS)

	fmt.Printf("Server started at port%s\n", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
