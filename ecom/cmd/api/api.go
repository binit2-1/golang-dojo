package api

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/binit2-1/golang-dojo/rest-api/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct{
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer{
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error{
	router := mux.NewRouter() //If you want to change the router, you can do it here
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}