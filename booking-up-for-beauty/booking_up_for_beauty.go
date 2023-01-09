package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/_2/2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January _2, 2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}

	now := time.Now()

	return t.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January _2, 2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}

	h := t.Hour()

	return h >= 12 && h < 18

}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/_2/2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}

	desc := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		t.Weekday(),
		t.Month(),
		t.Day(),
		t.Year(),
		t.Hour(),
		t.Minute(),
	)

	return desc
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// The beauty salon opened on September 15th in 2012.

	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
