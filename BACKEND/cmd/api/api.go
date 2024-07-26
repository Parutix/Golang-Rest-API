package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	address string
	db      *sql.DB
}

// Uppercase C in CreateNewAPIServer means that this function is exported
func CreateNewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	log.Println(subrouter)
	// Remove subrouter print line
	log.Println("Starting server on", s.address)
	return http.ListenAndServe(s.address, router)
}