package criptoHelper

import (
	"context"

	"github.com/cuwand/pondasi/observability"
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
	ctx, span := observability.Start(ctx, "compareHashAndPasswordCriptoHelper")
	defer span.End()

	return bcrypt.CompareHashAndPassword([]byte(hashedVal), []byte(val))
}
