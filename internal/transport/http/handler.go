// The package for our handler, running on the transport layer
package http

import (
	"github.com/gorilla/mux"
)

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
