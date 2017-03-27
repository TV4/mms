package titleservice

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	for _, tt := range []struct {
		hour   int
		minute int
		want   string
	}{
		{23, 45, "2345"},
		{0, 15, "2415"},
		{1, 45, "2545"},
		{2, 0, "0200"},
	} {
		if got := Time(tt.hour, tt.minute); got != tt.want {
			t.Fatalf("Time(%d, %d) = %q, want %q", tt.hour, tt.minute, got, tt.want)
		}
	}
}

func TestDate(t *testing.T) {
	for _, tt := range []struct {
		year  int
		month time.Month
		day   int
		want  string
	}{
		{0, 0, 0, ""},
		{0, 1, 0, ""},
		{0, 0, 2, ""},
		{2002, 12, 5, "20021205"},
		{2007, 6, 15, "20070615"},
		{2017, 8, 15, "20170815"},
		{2017, 3, 27, "20170327"},
	} {
		if got := Date(tt.year, tt.month, tt.day); got != tt.want {
			t.Fatalf("Date(%d, %d, %d) = %q, want %q", tt.year, tt.month, tt.day, got, tt.want)
		}
	}
}
