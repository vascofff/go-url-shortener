package handler

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandlers(t *testing.T) {
	v := validator.New()

	urlCreationRequest1 := UrlCreationRequest{
		Url: "",
	}
	err1 := v.Struct(urlCreationRequest1)

	urlCreationRequest2 := UrlCreationRequest{
		Url: "http://",
	}
	err2 := v.Struct(urlCreationRequest2)

	urlCreationRequest3 := UrlCreationRequest{
		Url:       "http://",
		ExpiresAt: "",
	}
	err3 := v.Struct(urlCreationRequest3)

	urlCreationRequest4 := UrlCreationRequest{
		Url:       "http://",
		ExpiresAt: "abvgde",
	}
	err4 := v.Struct(urlCreationRequest4)

	urlCreationRequest5 := UrlCreationRequest{
		Url:       "http://",
		ExpiresAt: "2006-01-02",
	}
	err5 := v.Struct(urlCreationRequest5)

	assert.NotEmpty(t, err1.Error())
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	assert.NotEmpty(t, err4.Error())
	assert.Nil(t, err5)
}
