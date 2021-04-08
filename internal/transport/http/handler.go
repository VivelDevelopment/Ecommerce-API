// The package for our handler, running on the transport layer
package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupHandler - sets up the handler
func SetupHandler() error {
	log.Println("Starting API handler on :8080")
	if err := http.ListenAndServe(":8080", Handle()); err != nil {
		log.Fatal("Failed to start up API service")
		return err
	}
	return nil
}

// Handler - returns a pointer to a mux router
func Handle() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// All of the routes
	router.HandleFunc("/", HomePage)
	router.HandleFunc("/api/v1/item/add", CreateItem).Methods("POST")
	router.HandleFunc("/api/v1/items/all", GetAllItems).Methods("GET")
	router.HandleFunc("/api/v1/item/{id}", GetOneItem).Methods("GET")
	router.HandleFunc("/api/v1/item/{id}", UpdateItem).Methods("PATCH")
	router.HandleFunc("/api/v1/item/{id}", DeleteItem).Methods("DELETE")
	return router
}
