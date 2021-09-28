package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/vascofff/go-url-shortener/src/db"
	"github.com/vascofff/go-url-shortener/src/response"
	"github.com/vascofff/go-url-shortener/src/shortener"
)

type UrlCreationRequest struct {
	Url       string `json:"url" validate:"required,url"`
	ExpiresAt string `json:"expires_at" validate:"omitempty,datetime=2006-01-02"`
}

func SendAGreeting(w http.ResponseWriter, _ *http.Request) {
	data := response.Data{"message": "Welcome to url shortener API"}
	response.JsonResponse(w, http.StatusOK, data)
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var creationRequest UrlCreationRequest
	err := json.NewDecoder(r.Body).Decode(&creationRequest)
	if err != nil {
		data := response.Data{"message": "Can't get data from request"}
		response.JsonResponse(w, http.StatusBadRequest, data)
		return
	}

	v := validator.New()
	err = v.Struct(creationRequest)
	if err != nil {
		data := response.Data{"message": err.Error()}
		response.JsonResponse(w, http.StatusBadRequest, data)
		return
	}

	err = creationRequest.UrlCreationRequestValidate()
	if err != nil {
		data := response.Data{"message": err.Error()}
		response.JsonResponse(w, http.StatusBadRequest, data)
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.Url)
	uuid := uuid.New().String()

	saveUrlMappingErr := db.SaveUrlMapping(uuid, shortUrl, creationRequest.Url, creationRequest.ExpiresAt)
	if saveUrlMappingErr != nil {
		data := response.Data{"message": saveUrlMappingErr.Error()}
		response.JsonResponse(w, http.StatusInsufficientStorage, data)
		return
	}

	host := "http://localhost:9808/"

	data := response.Data{"message": host + "long-url/" + uuid}
	response.JsonResponse(w, http.StatusCreated, data)
}

func HandleShortUrlRedirect(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	if uuid == "" {
		data := response.Data{"message": "Uuid can't be empty string"}
		response.JsonResponse(w, http.StatusBadRequest, data)
		return
	}

	initialUrl, expiresAt, retrieveInitialUrlErr := db.RetrieveInitialUrl(uuid)
	if retrieveInitialUrlErr != nil {
		data := response.Data{"message": retrieveInitialUrlErr.Error()}
		response.JsonResponse(w, http.StatusInsufficientStorage, data)
		return
	}

	if expiresAt != "" {
		if isDateExpired(expiresAt) {
			data := response.Data{"message": "Date for given uuid is already expired"}
			response.JsonResponse(w, http.StatusCreated, data)
			return
		}
	}

	http.Redirect(w, r, initialUrl, http.StatusMovedPermanently)
}
