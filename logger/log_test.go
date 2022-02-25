package logger

import (
	"testing"
)

func TestLog(t *testing.T) {

	GetAppLogger().Error("Hello")
	GetAppLogger().Info("Hellees")

}
