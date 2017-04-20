package titleservice

import (
	"fmt"
	"time"
)

// Stockholm is the Time Zone in Sweden
var Stockholm *time.Location

func init() {
	if location, err := time.LoadLocation("Europe/Stockholm"); err == nil {
		Stockholm = location
	}
}

// Time formats a hour and minute into the MMS Time format
//
// 0200-2559 (The leading zero is optional)
//
// Examples
//
//   Time(23,45) = 2345 (day 1)
//   Time(0,15)  = 2415 (day 1)
//   Time(1,45)  = 2545 (day 1)
//   Time(2,0)   = 0200 (day 2)
//
func Time(hour, minute int) string {
	if hour < 2 {
		return fmt.Sprintf("%02d%02d", hour+24, minute)
	}

	return fmt.Sprintf("%02d%02d", hour, minute)
}

// Date formats a year, month, day into the format YYYYMMDD
func Date(year int, month time.Month, day int) string {
	if year > 0 && month > 0 && month < 13 && day > 0 && day < 32 {
		return fmt.Sprintf("%04d%02d%02d", year, month, day)
	}

	return ""
}

// DateAtTime returns the date string for the provided time.Time
func DateAtTime(t time.Time) string {
	return Date(t.In(Stockholm).Add(-2 * time.Hour).Date())
}
