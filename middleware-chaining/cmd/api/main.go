package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)




func main(){
	//new router
	mux := mux.NewRouter()

	//health route
	mux.HandleFunc("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"messages:Health OK...}`))
	} )

	//port
	port := ":8080"
	fmt.Printf("Starting Server on port %s\n", port)
	
	err := http.ListenAndServe(port, mux)
	if err != nil{
		log.Fatalf("Failed to start server: %v", err)
	}

}