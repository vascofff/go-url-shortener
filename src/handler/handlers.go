package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/vascofff/go-url-shortener/src/response"
	"github.com/vascofff/go-url-shortener/src/url"
	"github.com/vascofff/go-url-shortener/validation"
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

	urlValidateErr := validation.ValidateUrl(creationRequest.Url)
	if urlValidateErr != nil {
		data := response.Data{"message": urlValidateErr.Error()}
		response.JsonResponse(w, http.StatusBadRequest, data)
		return
	}

	_, expiresAtValidateErr := validation.ValidateExpiresAt(&creationRequest.ExpiresAt)
	if expiresAtValidateErr != nil {
		data := response.Data{"message": expiresAtValidateErr.Error()}
		response.JsonResponse(w, http.StatusBadRequest, data)
		return
	}

	uuid, createErr := url.CreateShortUrl(creationRequest.Url, &creationRequest.ExpiresAt)
	if createErr != nil {
		data := response.Data{"message": createErr.Error()}
		response.JsonResponse(w, http.StatusInternalServerError, data)
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

	initialUrl, expiresAt, err := url.GetLongUrlWithExpiresAt(uuid)
	if err != nil {
		data := response.Data{"message": err.Error()}
		response.JsonResponse(w, http.StatusInternalServerError, data)
		return
	}

	if expiresAt != nil {
		if validation.IsDateExpired(*expiresAt) {
			data := response.Data{"message": err.Error()}
			response.JsonResponse(w, http.StatusOK, data)
			return
		}
	}

	http.Redirect(w, r, initialUrl, http.StatusMovedPermanently)
}
