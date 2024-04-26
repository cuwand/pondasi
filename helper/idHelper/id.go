package idHelper

import (
	"github.com/oklog/ulid/v2"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func UUID() string {
	u4 := uuid.NewV4()
	return u4.String()
}

func ULID() string {
	return ulid.Make().String()
}

func UUIDClean() string {
	u4 := uuid.NewV4()
	return strings.ReplaceAll(u4.String(), "-", "")
}

func ReferenceNumber() string {
	//return strings.ToUpper(stringHelper.RandomAlphaNumeric(30))
	return strings.ToUpper(UUIDClean())
}
