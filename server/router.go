package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//defining routes
func Routes() {
	routes := mux.NewRouter()
	s := routes.PathPrefix("/api").Subrouter()

	log.Fatal(http.ListenAndServe(":8080", s))
}