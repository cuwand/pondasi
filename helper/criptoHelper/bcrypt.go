package criptoHelper

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateBcrypt(val string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func CompareHashAndPassword(hashedVal, val string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedVal), []byte(val))
}
