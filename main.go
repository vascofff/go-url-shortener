package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/vascofff/go-url-shortener/src/db"
	"github.com/vascofff/go-url-shortener/src/handler"
)

func main() {
	// If the file doesn't exist, create it or append to the file
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	//     log.Fatal(err)
	// }

	// log.SetOutput(file)

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

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Welcome to url shortener API")
	})

	r.Post("/create-short-url", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateShortUrl(w, r)
	})

	r.Get("/{uuid:[a-zA-Z0-9-]+}", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleShortUrlRedirect(w, r)
	})

	http.ListenAndServe(":9808", r)
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	// }
}
