package common

import (
	"math/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"
)

const randomLength  =16

func GenerateSalt(length int) string {
	var salt []byte
	var asciiPad int64

	if 0 == length {
		length = randomLength
	}
	
	asciiPad = 32
	for i := 0; i < length; i++ {
		salt = append(salt , byte(rand.Int63n(94) + asciiPad))
	}

	return string(salt)
}

func GenerateHash(salt string ,password string) string  {
	var hash string
	fullstring := salt + password
	sha := sha256.New()
	sha.Write([]byte(fullstring))
	hash = base64.URLEncoding.EncodeToString(sha.Sum(nil))

	return hash
}

func ReturnPassword(password string) (string ,string)  {
	rand.Seed(time.Now().UTC().UnixNano())

	salt := GenerateSalt(0)
	hash := GenerateHash(salt , password)

	return salt, hash
}

