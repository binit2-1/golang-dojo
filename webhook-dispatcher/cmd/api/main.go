package main

import (
	"net/http"

	"github.com/binit2-1/golang-dojo/webhook-dispatcher/internal/server"
)



func main(){

	mux := http.NewServeMux()

	port := ":8080"

	mux.HandleFunc("POST /v1/webhooks/fire", server.FireWebhookHandler)



	println("Starting Server on port", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		println("Failed to start server:", err.Error())
	}
}