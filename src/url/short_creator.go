package url

import (
	"github.com/google/uuid"
	"github.com/vascofff/go-url-shortener/src/db"
	"github.com/vascofff/go-url-shortener/src/shortener"
)

func CreateShortUrl(url string, expiresAt *string) (string, error) {
	shortUrl, shortLinkGenErr := shortener.GenerateShortLink(url)
	if shortLinkGenErr != nil {
		return "", shortLinkGenErr
	}

	uuid := uuid.New().String()

	saveUrlMappingErr := db.SaveUrlMapping(uuid, shortUrl, url, expiresAt)
	if saveUrlMappingErr != nil {
		return "", saveUrlMappingErr
	}

	return uuid, nil
}
