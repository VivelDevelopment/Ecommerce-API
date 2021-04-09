package main

import (
	transportDB "github.com/VivelDevelopment/Ecommerce-API/internal/database"
	transportHTTP "github.com/VivelDevelopment/Ecommerce-API/internal/transport/http"
)

// An empty struct for our service
type Service struct{}

// Method to start our service
func (s *Service) Start() {
	transportHTTP.SetupHandler()
	transportDB.SetupDB()
}

// Entry point
func main() {
	var service Service
	service.Start()
}
