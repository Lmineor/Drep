package utils

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
	t, _ := time.Parse("2006-01-02", timeStr)
	// Because Bj TimeZone is +8, so we sbutract 8hours
	zeroTime := t.Unix() - 28800
	return zeroTime
}