package model

import (
	"database/sql/driver"
	"time"
)

const DateFormat = "2006-01-02"

type LocalDate time.Time

func (t *LocalDate) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalDate(time.Time{})
		return
	}

	location, _ := time.LoadLocation(Loc)

	now, err := time.ParseInLocation(`"`+DateFormat+`"`, string(data), location)
	*t = LocalDate(now)
	return
}

func (t LocalDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(DateFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, DateFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalDate) Value() (driver.Value, error) {
	if t.String() == "0001-01-01" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(DateFormat)), nil
}

func (t *LocalDate) Scan(v interface{}) error {
	location, _ := time.LoadLocation(Loc)
	tTime, _ := time.ParseInLocation("2006-01-02 +0800 CST", v.(time.Time).String(), location)
	*t = LocalDate(tTime)
	return nil
}

func (t LocalDate) String() string {
	return time.Time(t).Format(DateFormat)
}
