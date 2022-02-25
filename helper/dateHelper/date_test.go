package dateHelper

import (
	"github.com/cuwand/pondasi/logger"
	"os"
	"testing"
)

func init() {
	os.Setenv("TZ", "Asia/Jakarta")
}

func TestDateTimeNow(t *testing.T) {

	logger.GetAppLogger().Info(TimeNow().Format(DateTimeFormat))
	logger.GetAppLogger().Info(BeginingOfDay().Format(DateTimeFormat))
	logger.GetAppLogger().Info(EndOfDay().Format(DateTimeFormat))

	logger.GetAppLogger().Info(BeginingOfMonth().Format(DateFormat))
	logger.GetAppLogger().Info(EndOfMonth().Format(DateFormat))

}
