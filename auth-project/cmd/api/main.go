package main

import (
	"fmt"
	"log"
	"net/http"

)

func main(){

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"messages":"auth-project up"}`))
	})


	port := ":8080"
	fmt.Printf("Starting Server on port %s\n", port)


	


	err := http.ListenAndServe(port, mux)
	if err != nil{
		log.Fatalf("Failed to start server: %v", err)
	}

	
}
