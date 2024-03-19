package model

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

const Loc = "Asia/Shanghai"

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	location, _ := time.LoadLocation(Loc)

	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), location)
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	location, _ := time.LoadLocation(Loc)
	tTime, _ := time.ParseInLocation("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String(), location)
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}
