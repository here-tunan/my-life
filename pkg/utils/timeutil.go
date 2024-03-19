package utils

import (
	"fmt"
	"go-my-life/pkg/model"
	"time"
)

const formatter = "2006-01-02 15:04:05"
const Loc = "Asia/Shanghai"

func Format(t time.Time) string {
	location, err := time.LoadLocation(Loc)
	if err != nil {
		fmt.Println("时区加载错误:", err)
		return ""
	}
	// 转换为指定时区的时间
	t = t.In(location)

	formattedTime := t.Format(formatter)

	return formattedTime
}

func LocalTimeFormat(lt model.LocalTime) string {
	t := time.Time(lt)
	location, err := time.LoadLocation(Loc)
	if err != nil {
		fmt.Println("时区加载错误:", err)
		return ""
	}
	// 转换为指定时区的时间
	t = t.In(location)

	formattedTime := t.Format(formatter)

	return formattedTime
}

func ParseTimeInLoc(formatter string, timeStr string) time.Time {
	location, err := time.LoadLocation(Loc)
	if err != nil {
		fmt.Println("时区加载错误:", err)
		return time.Time{}
	}
	curTime, err := time.ParseInLocation(formatter, timeStr, location)
	return curTime
}
