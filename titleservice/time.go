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

// Time formats an hour and minute into the format HHMM
func Time(hour, minute int) string {
	return fmt.Sprintf("%02d%02d", hour, minute)
}

// BroadcastTime formats an hour and minute into the MMS Live TV broadcast time format
//
// 0200-2559 (The leading zero is optional)
//
// Examples
//
//   BroadcastTime(23,45) = 2345 (day 1)
//   BroadcastTime(0,15)  = 2415 (day 1)
//   BroadcastTime(1,45)  = 2545 (day 1)
//   BroadcastTime(2,0)   = 0200 (day 2)
//
func BroadcastTime(hour, minute int) string {
	if hour < 2 {
		return Time(hour+24, minute)
	}

	return Time(hour, minute)
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
	return Date(t.In(Stockholm).Date())
}

// BroadcastDateAtTime returns the Live TV broadcast date string for the provided time.Time
func BroadcastDateAtTime(t time.Time) string {
	return DateAtTime(t.Add(-2 * time.Hour))
}
