package main

import (
	"fmt"
	"log"
	"net/http"

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

	//port
	port := ":8080"
	fmt.Printf("Starting Server on port %s\n", port)
	
	err := http.ListenAndServe(port, mux)
	if err != nil{
		log.Fatalf("Failed to start server: %v", err)
	}

}