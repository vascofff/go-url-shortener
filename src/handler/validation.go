package handler

import (
	"net/url"
	"regexp"
	"time"

	"github.com/pkg/errors"
)

const DateLayout = "2006-01-02"

func (urlCreationRequest *UrlCreationRequest) UrlCreationRequestValidate() error {
	err := urlCreationRequest.validateUrl()
	if err != nil {
		return errors.Wrap(err, err.Error())
	}

	err = urlCreationRequest.validateExpiresAt()
	if err != nil {
		return errors.Wrap(err, err.Error())
	}

	return nil
}

func (urlCreationRequest UrlCreationRequest) validateUrl() error {
	_, err := url.ParseRequestURI(urlCreationRequest.Url)

	if err != nil {
		return errors.New("not correct url format. Must starts with http(s)://")
	}

	return nil
}

func (urlCreationRequest *UrlCreationRequest) validateExpiresAt() error {
	if urlCreationRequest.ExpiresAt != "" {
		expiresAt, err := time.Parse(DateLayout, urlCreationRequest.ExpiresAt)
		if err != nil {
			return errors.New("error while parsing input expires_at data")
		}
		if expiresAt.In(time.UTC).Before(time.Now().In(time.UTC)) {
			return errors.New("the received expires_at should not be earlier than tomorrow")
		}
	}

	return nil
}

func isDateExpired(expiresAt string) bool {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	parsedExpiresAt, _ := time.Parse(DateLayout, re.FindString(expiresAt))

	return parsedExpiresAt.In(time.UTC).Before(time.Now().In(time.UTC))
}
