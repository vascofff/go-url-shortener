package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/vascofff/go-url-shortener/src/db"
	"github.com/vascofff/go-url-shortener/src/shortener"
)

type UrlCreationRequest struct {
	Url       string `json:"url"`
	ExpiresOn string `json:"expires_on"`
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var creationRequest UrlCreationRequest
	json.NewDecoder(r.Body).Decode(&creationRequest)

	creationRequest.UrlCreationRequestValidate()

	shortUrl := shortener.GenerateShortLink(creationRequest.Url)
	uuid := uuid.New().String()
	db.SaveUrlMapping(uuid, shortUrl, creationRequest.Url, creationRequest.ExpiresOn)

	host := "http://localhost:9808/"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(host + uuid)
}

func HandleShortUrlRedirect(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	if uuid == "" {
		log.Fatalf("Received uuid is empty")
	}

	initialUrl, expiresOn := db.RetrieveInitialUrl(uuid)

	if isDateExpired(expiresOn) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Date for given uuid is already expired")
	}

	http.Redirect(w, r, initialUrl, 302)
}
