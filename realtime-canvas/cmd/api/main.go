package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"realtime-canvas/internal/server"
	"realtime-canvas/internal/websocket"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()

	hub := websocket.NewHub()

	go hub.Run()

	port := os.Getenv("PORT")
	
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.ServeWS(hub, w, r)
	})

	fmt.Printf("Server started at port%s\n", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
