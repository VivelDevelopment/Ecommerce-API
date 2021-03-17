package handler

import (
	"log"

	"github.com/VivelDevelopment/Ecommerce-API/api"
	"github.com/gorilla/mux"
)

func Handle() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	if router != nil {
		log.Println("Instance started on :8080")
	}
	router.HandleFunc("/", api.HomePage)
	router.HandleFunc("/api/v1/item/add", api.CreateItem).Methods("POST")
	router.HandleFunc("/api/v1/items/all", api.GetAllItems).Methods("GET")
	router.HandleFunc("/api/v1/item/{id}", api.GetOneItem).Methods("GET")
	router.HandleFunc("/api/v1/item/{id}", api.UpdateItem).Methods("PATCH")
	router.HandleFunc("/api/v1/item/{id}", api.DeleteItem).Methods("DELETE")
	return router
}
