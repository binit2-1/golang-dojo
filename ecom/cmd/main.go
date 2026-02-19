package main

import (
	"github.com/binit2-1/golang-dojo/rest-api/cmd/api"
	"log"
)



func main(){
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil{
		log.Fatal(err)
	}
}