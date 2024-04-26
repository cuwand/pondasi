package hmac

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

func Hashed(key, plainData string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	//h := hmac.New(sha256.New, []byte(key))
	h := hmac.New(sha512.New, []byte(key))

	// Write Data to it
	h.Write([]byte(plainData))

	// Get result and encode as hexadecimal string
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
