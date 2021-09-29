package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
	"github.com/pkg/errors"
)

func sha256Of(input string) []byte {
	algorith := sha256.New()
	algorith.Write([]byte(input))
	return algorith.Sum(nil)
}

func base58Encoded(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		return "", errors.New("failed to generate short link")
	}
	return string(encoded), nil
}

func GenerateShortLink(initialLink string) (string, error) {
	urlHashBytes := sha256Of(initialLink)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString, err := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	if err != nil {
		return "", errors.Wrap(err, err.Error())
	}

	return finalString[:8], nil
}
