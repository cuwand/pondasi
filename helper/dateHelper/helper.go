package dateHelper

import (
	"time"

	"github.com/cuwand/pondasi/helper/envHelper"
)

func Loc() *time.Location {
	locName := envHelper.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return loc
}

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

func ToLocalTimeZone(val time.Time) time.Time {
	locName := envHelper.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return val.In(loc)
}

// time database to Asia/Jakarta
func TimeToString(dateTime time.Time) string {
	locName := envHelper.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	return dateTime.In(loc).Format(DateTimeFormat)
}

func StringToTime(dateTimeStr string) time.Time {
	dateTime, err := time.Parse(DateTimeFormat, dateTimeStr)

	if err != nil {
		panic(err)
	}

	return dateTime
}

func StringToTimeFormat(dateTimeStr string, format string) time.Time {
	locName := envHelper.GetEnvAndValidate("TZ")
	loc, _ := time.LoadLocation(locName)

	dateTime, err := time.ParseInLocation(format, dateTimeStr, loc)

	if err != nil {
		panic(err)
	}

	return dateTime
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

func ToBeginingOfDay(dateTime time.Time) time.Time {
	y, m, d := dateTime.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, dateTime.Location())
}

func ToEndOfDay(dateTime time.Time) time.Time {
	y, m, d := dateTime.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), dateTime.Location())
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
