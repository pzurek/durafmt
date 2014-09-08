package durafmt

import (
	"fmt"
	"time"
)

func HMS(duration time.Duration) string {

	h, m, s := extractValues(duration)

	hours := fmt.Sprintf("%0.2d", h)
	minutes := fmt.Sprintf("%0.2d", m)
	seconds := fmt.Sprintf("%0.2d", s)

	d := fmt.Sprintf("%s:%s:%s", hours, minutes, seconds)

	return d
}

func extractValues(duration time.Duration) (int64, int64, int64) {
	var hours, minutes, seconds int64
	if duration.Hours() < 1 {
		hours = 0
	} else {
		hours = int64(duration.Hours())
	}

	if duration.Minutes() < 1 {
		minutes = 0
	} else {
		minutes = int64(duration.Minutes()) - hours*60
	}

	if duration.Seconds() < 1 {
		seconds = 0
	} else {
		seconds = int64(duration.Seconds()) - hours*60*60 - minutes*60
	}

	return hours, minutes, seconds
}
