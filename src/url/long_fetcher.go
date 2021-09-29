package url

import (
	"time"

	"github.com/pkg/errors"
	"github.com/vascofff/go-url-shortener/src/db"
	"github.com/vascofff/go-url-shortener/validation"
)

func GetLongUrlWithExpiresAt(uuid string) (string, *time.Time, error) {
	initialUrl, expiresAt, retrieveInitialUrlErr := db.RetrieveInitialUrl(uuid)
	if retrieveInitialUrlErr != nil {
		return "", nil, errors.Wrap(retrieveInitialUrlErr, retrieveInitialUrlErr.Error())
	}

	parsedExpiresAt, _ := time.Parse(validation.DateLayout, expiresAt)

	return initialUrl, &parsedExpiresAt, nil
}
