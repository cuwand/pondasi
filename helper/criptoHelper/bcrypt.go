package criptoHelper

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func GenerateBcrypt(val string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func CompareHashAndPassword(ctx context.Context, hashedVal, val string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedVal), []byte(val))
}
