package stringHelper

import (
	"math/rand"
)

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func randomStringSet(n int, set string) (string, error) {
	bytes, err := randomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = set[b%byte(len(set))]
	}
	return string(bytes), nil
}

func RandomAlphaNumeric(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	return randomStringSet(n, letters)
}

func RandomNumericString(n int) (string, error) {
	const letters = "0123456789"
	return randomStringSet(n, letters)
}