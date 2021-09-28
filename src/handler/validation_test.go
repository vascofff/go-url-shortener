package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDateExpiredAt(t *testing.T) {
	expiredDate1 := "2000-00-00"
	isDateExpired1 := isDateExpired(expiredDate1)

	expiredDate2 := "2021-09-09"
	isDateExpired2 := isDateExpired(expiredDate2)

	notExpiredDate1 := "2055-01-02"
	isDateNotExpired1 := isDateExpired(notExpiredDate1)

	assert.Equal(t, isDateExpired1, true)
	assert.Equal(t, isDateExpired2, true)
	assert.Equal(t, isDateNotExpired1, false)
}
