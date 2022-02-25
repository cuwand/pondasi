package dateHelper

import (
	"github.com/cuwand/pondasi/helper/envHelper"
	"time"
)

func TimeNow() time.Time {
	locName := envHelper.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return time.Now().In(loc)
}

func TimeNowUTC() time.Time {
	locName := envHelper.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return time.Now().In(loc).UTC()
}

func BeginingOfDay() time.Time {
	timeNow := TimeNow()

	y, m, d := timeNow.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, timeNow.Location())
}

func EndOfDay() time.Time {
	timeNow := TimeNow()

	y, m, d := timeNow.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), timeNow.Location())
}

func BeginingOfMonth() time.Time {
	timeNow := TimeNow()
	y, m, _ := timeNow.Date()

	return time.Date(y, m, 1, 0, 0, 0, 0, timeNow.Location())
}

func EndOfMonth() time.Time {
	timeNow := TimeNow()
	y, m, _ := timeNow.Date()

	return time.Date(y, m+1, 1, 0, 0, 0, -1, timeNow.Location())
}
