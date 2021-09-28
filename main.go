package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vascofff/go-url-shortener/router"
	"github.com/vascofff/go-url-shortener/src/db"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := db.Initialize(dbUser, dbPassword, dbName)

	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	defer database.Conn.Close()

	router := router.RegisterRoutes()

	err = http.ListenAndServe(":9808", router)
	if err != nil {
		log.Fatalf("Failed to start the web server - Error: %v", err)
	}
}
