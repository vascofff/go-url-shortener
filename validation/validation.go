package validation

import (
	"net/url"
	"time"

	"github.com/pkg/errors"
)

const DateLayout = "2006-01-02"

func ValidateUrl(requestUrl string) error {
	_, err := url.ParseRequestURI(requestUrl)

	if err != nil {
		return errors.New("not correct url format. Must starts with http(s)://")
	}

	return nil
}

func ValidateExpiresAt(requestExpiresAt *string) (*time.Time, error) {
	if requestExpiresAt != nil {
		parsedDate, err := time.Parse(DateLayout, *requestExpiresAt)
		if err != nil {
			return nil, errors.New("error while parsing input expires_at data. Send expires_at in format: 'YYYY-MM-DD'")
		}
		if parsedDate.In(time.UTC).Before(time.Now().In(time.UTC)) {
			return nil, errors.New("the received expires_at should not be earlier than tomorrow")
		}

		return &parsedDate, nil
	}

	return nil, nil
}

func IsDateExpired(expiresAt time.Time) bool {
	return expiresAt.In(time.UTC).Before(time.Now().In(time.UTC))
}
