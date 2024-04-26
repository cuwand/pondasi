package dateProperty

import (
	"fmt"
	"github.com/cuwand/pondasi/helper/dateHelper"
	"strings"
	"time"
)

var nilTime = (time.Time{}).UnixNano()

type Date struct {
	time.Time
}

func (ct *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	ct.Time, err = time.ParseInLocation(dateHelper.DateFormat, s, loc)
	return
}

func (ct *Date) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(dateHelper.DateFormat))), nil
}

func (ct *Date) ToDateFormat() string {
	return ct.Time.Format(dateHelper.DateFormat)
}

func (ct *Date) ToDateISOFormat() string {
	return ct.Time.Format(dateHelper.DateISOFormat)
}

func StringToDate(val string) Date {
	date, err := time.Parse(dateHelper.DateFormat, val)

	if err != nil {
		panic(fmt.Sprintf("Format must be %s", dateHelper.DateFormat))
	}

	return Date{
		date,
	}
}
