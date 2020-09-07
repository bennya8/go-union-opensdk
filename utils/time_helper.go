package utils

import (
	"strings"
	"time"
)

// 2006-01-02T15:04:05Z07:00
const (
	TimeLayoutYear       = "2006"
	TimeLayoutYearSimple = "06"
	TimeLayoutMonth      = "01"
	TimeLayoutDay        = "02"
	TimeLayoutHour       = "15"
	TimeLayoutMinute     = "04"
	TimeLayoutSecond     = "05"
)

func TimeDateToTimestamp(layout string, snapshot string) (int64, error) {
	t, err := time.Parse(layout, snapshot)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func TimeFormatToLayout(format string) (s string) {
	tplMaps := map[string]string{
		"YYYY": TimeLayoutYear,
		"YY":   TimeLayoutYearSimple,
		"MM":   TimeLayoutMonth,
		"DD":   TimeLayoutDay,
		"HH":   TimeLayoutHour,
		"mm":   TimeLayoutMinute,
		"ss":   TimeLayoutSecond,
	}

	for k, v := range tplMaps {
		format = strings.Replace(format, k, v, -1)
	}
	return format
}

func TimeFormatToNowString(format ...string) (s string) {
	layout := "YYYY-MM-DD HH:mm:ss"
	if len(format) > 0 {
		layout = format[0]
	}
	return time.Now().Format(TimeFormatToLayout(layout))
}
