package main

import (
	"log"
	"net/http"

	transportHTTP "github.com/VivelDevelopment/Ecommerce-API/internal/transport/http"
)

// An empty struct for our service
type Service struct{}

// Method to start our service
func (s *Service) Start() error {
	log.Println("Starting API service on :8080")
	if err := http.ListenAndServe(":8080", transportHTTP.Handle()); err != nil {
		log.Fatal("Failed to start up API service")
		return err
	}
	return nil
}

// Entry point
func main() {
	var service Service
	service.Start()
}
