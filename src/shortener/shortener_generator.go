package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
	// 	"math/rand"
	// 	"time"
)

// var (
// 	randomPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	poolLength = len(randomPool)
// )

func sha256Of(input string) []byte {
	algorith := sha256.New()
	algorith.Write([]byte(input))
	return algorith.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink string) string {
	urlHashBytes := sha256Of(initialLink)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

// func GenerateShortLink() string {
// 	return RandomString(8)
// }

// func RandomString(length int) string {
// 	str := make([]byte, length)
// 	rand.Seed(time.Now().UTC().UnixNano())

// 	for i := range str {
// 		str[i] = randomPool[rand.Intn(poolLength)]
// 	}

// 	return string(str)
// }
