package main

import (
	"log"
	"net/http"

	"github.com/VivelDevelopment/Ecommerce-API/handler"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", handler.Handle()))
}
