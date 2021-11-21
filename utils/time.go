package utils

import (
	"fmt"
	"time"
)

func InOneDay(timeStr string) bool {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	unixTime := t.Unix()
	midNightTime := getMidNight0UnixTime()
	if unixTime < midNightTime {
		return false
	} else {
		return (time.Now().Unix() - unixTime) < 86400
	}
}

func getMidNight0UnixTime() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	zeroTime := t.Unix()
	return zeroTime
}

func GetDateTimeStr() string {
	timeNow := time.Now()
	year, month, day := timeNow.Date()
	dateStr := fmt.Sprintf("%d-%d-%d", year, month, day)
	hour, min, sec := timeNow.Clock()
	timeStr := fmt.Sprintf("%d%d%d", hour, min, sec)
	return fmt.Sprintf("%s-%s", dateStr, timeStr)
}
