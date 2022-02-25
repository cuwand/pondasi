package idHelper

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func UUID() string {
	u4 := uuid.NewV4()
	return u4.String()
}

func UUIDClean() string {
	u4 := uuid.NewV4()
	return strings.ReplaceAll(u4.String(), "-", "")
}