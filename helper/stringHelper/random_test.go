package stringHelper

import (
	"github.com/cuwand/pondasi/helper/idHelper"
	"github.com/cuwand/pondasi/helper/slugHelper"
	"github.com/cuwand/pondasi/logger"
	"strings"
	"testing"
)

func TestStringHelper(t *testing.T)  {

	randomAlphaNumeric, err := RandomAlphaNumeric(23)

	if err != nil {
		panic(err)
	}

	logger.GetAppLogger().Info(randomAlphaNumeric)

	randomNumericString, err := RandomNumericString(23)

	if err != nil {
		panic(err)
	}

	logger.GetAppLogger().Info(randomNumericString)

	logger.GetAppLogger().Info(idHelper.UUID())
	logger.GetAppLogger().Info(strings.ToUpper(idHelper.UUIDClean()))
	logger.GetAppLogger().Info(slugHelper.Generate("Lowercase defines if the resulting slug is transformed to lowercase"))

}