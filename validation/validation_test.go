package validation

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsDateExpiredAt(t *testing.T) {
	expiredDate1 := "2000-00-00"
	parsedDate1, _ := time.Parse(DateLayout, expiredDate1)
	isDateExpired1 := IsDateExpired(parsedDate1)

	expiredDate2 := "2021-09-09"
	parsedDate2, _ := time.Parse(DateLayout, expiredDate2)
	isDateExpired2 := IsDateExpired(parsedDate2)

	notExpiredDate1 := "2055-01-02"
	parsedDate3, _ := time.Parse(DateLayout, notExpiredDate1)
	isDateNotExpired1 := IsDateExpired(parsedDate3)

	assert.Equal(t, isDateExpired1, true)
	assert.Equal(t, isDateExpired2, true)
	assert.Equal(t, isDateNotExpired1, false)
}
