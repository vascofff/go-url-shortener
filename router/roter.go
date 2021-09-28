package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/vascofff/go-url-shortener/src/handler"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", handler.SendAGreeting)
	r.MethodFunc("POST", "/create-short-url", handler.CreateShortUrl)
	r.MethodFunc("GET", "/{uuid:[a-zA-Z0-9-]+}", handler.HandleShortUrlRedirect)

	r.NotFoundHandler()

	return r
}
