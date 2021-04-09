package database

import (
	"fmt"
	"log"
	"os"

	transportHTTP "github.com/VivelDevelopment/Ecommerce-API/internal/transport/http"
	"github.com/jinzhu/gorm"
)

var DBConn *gorm.DB

// SetupDB - sets up the database
func SetupDB() {
	var (
		err    error
		dbHost = os.Getenv("DB_HOST")
		dbUser = os.Getenv("DB_USERNAME")
		dbName = os.Getenv("DB_NAME")
		dbPass = os.Getenv("DB_PASSWORD")
		dbPort = os.Getenv("DB_PORT")
		dbURI  = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, dbUser, dbName, dbPass, dbPort)
	)
	DBConn, err = gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("Could not start database")
	}
	DBConn.AutoMigrate(&transportHTTP.Item{})
}
