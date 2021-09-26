package handler

import (
	"log"
	"net/url"
	"regexp"
	"time"
)

const DateLayout = "2006-01-02"

func (urlCreationRequest *UrlCreationRequest) UrlCreationRequestValidate() {
	urlCreationRequest.validateUrl()
	urlCreationRequest.validateExpiresOn()
}

func (urlCreationRequest UrlCreationRequest) validateUrl() {
	if urlCreationRequest.Url == "" {
		log.Fatalf("url param can't be empty")
	}

	_, err := url.Parse(urlCreationRequest.Url)

	if err != nil {
		log.Fatal("Not correct url format")
	}
}

func (urlCreationRequest *UrlCreationRequest) validateExpiresOn() {
	var expiresOnRegexp string = `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`

	if urlCreationRequest.ExpiresOn == "" {
		expiresOn := time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC)
		urlCreationRequest.ExpiresOn = expiresOn.String()
		return
	}

	expRe := regexp.MustCompile(expiresOnRegexp)
	if !expRe.MatchString(urlCreationRequest.ExpiresOn) {
		log.Fatal("Incorrect expires_on format. Must be like YYYY-MM-DD")
	}

	expiresOn, _ := time.ParseInLocation(DateLayout, urlCreationRequest.ExpiresOn, time.UTC)
	if expiresOn.In(time.UTC).Before(time.Now().In(time.UTC)) {
		log.Fatalf("Given expires_on can't be date in past")
	}

	urlCreationRequest.ExpiresOn = expiresOn.String()

	return
}

func isDateExpired(expiresOn string) bool {
	parsedExpiresOn, _ := time.ParseInLocation(DateLayout, expiresOn, time.UTC)
	return parsedExpiresOn.In(time.UTC).Before(time.Now().In(time.UTC))
}
