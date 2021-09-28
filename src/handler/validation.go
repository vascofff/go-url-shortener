package handler

import (
	"errors"
	"net/url"
	"regexp"
	"time"
)

const DateLayout = "2006-01-02"

func (urlCreationRequest *UrlCreationRequest) UrlCreationRequestValidate() string {
	err := urlCreationRequest.validateUrl()
	if err != nil {
		return err.Error()
	}

	err = urlCreationRequest.validateExpiresAt()
	if err != nil {
		return err.Error()
	}

	return ""
}

func (urlCreationRequest UrlCreationRequest) validateUrl() error {
	_, err := url.ParseRequestURI(urlCreationRequest.Url)

	if err != nil {
		return errors.New("Not correct url format. Must starts with http(s)://")
	}

	return nil
}

func (urlCreationRequest *UrlCreationRequest) validateExpiresAt() error {
	if urlCreationRequest.ExpiresAt != "" {
		expiresAt, _ := time.Parse(DateLayout, urlCreationRequest.ExpiresAt)
		if expiresAt.In(time.UTC).Before(time.Now().In(time.UTC)) {
			return errors.New("The received expires_at should not be earlier than tomorrow")
		}
	}

	return nil
}

func isDateExpired(expiresAt string) bool {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	parsedExpiresAt, _ := time.Parse(DateLayout, re.FindString(expiresAt))

	return parsedExpiresAt.In(time.UTC).Before(time.Now().In(time.UTC))
}
