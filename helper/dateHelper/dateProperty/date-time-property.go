package dateProperty

import (
	"fmt"
	"github.com/cuwand/pondasi/helper/dateHelper"
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

func ToDateTime(val time.Time) DateTime {
	return DateTime{
		val,
	}
}

func (ct *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	ct.Time, err = time.ParseInLocation(dateHelper.DateTimeFormat, s, loc)
	return
}

func (ct *DateTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(dateHelper.DateTimeFormat))), nil
}

func (ct *DateTime) ToDateTimeFormat() string {
	return ct.Time.Format(dateHelper.DateTimeFormat)
}

func (ct *DateTime) ToDateTimeISOFormat() string {
	return ct.Time.Format(dateHelper.DateTimeISOFormat)
}

func (ct DateTime) BeginingOfDay() time.Time {
	y, m, d := ct.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, ct.Location())
}

func (ct DateTime) EndOfDay() time.Time {
	y, m, d := ct.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), ct.Location())
}

func StringToDateTime(val string) DateTime {
	date, err := time.Parse(dateHelper.DateFormat, val)

	if err != nil {
		panic(fmt.Sprintf("Format must be %s", dateHelper.DateFormat))
	}

	return DateTime{
		date,
	}
}
