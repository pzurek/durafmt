package durafmt

import (
	"fmt"
	"time"
)

// HMS function returns duration as a string in format 'hh:mm:ss'
func HMS(duration time.Duration) string {

	h, m, s := extractValues(duration)

	hours := fmt.Sprintf("%0.2d", h)
	minutes := fmt.Sprintf("%0.2d", m)
	seconds := fmt.Sprintf("%0.2d", s)

	d := fmt.Sprintf("%s:%s:%s", hours, minutes, seconds)

	return d
}

// LongWords function returns duration as a string in format 'X hours Y minutes Z seconds'
// Only non-zero values will be taken into account
func LongWords(duration time.Duration) string {

	var d string
	h, m, s := extractValues(duration)

	seconds := fmt.Sprintf("%d", s)
	d = fmt.Sprintf("%s seconds", seconds)

	if duration.Minutes() < 1 {
		return d
	}

	minutes := fmt.Sprintf("%d", m)
	d = fmt.Sprintf("%s minutes %s seconds", minutes, seconds)

	if duration.Hours() < 1 {
		return d
	}

	hours := fmt.Sprintf("%d", h)
	d = fmt.Sprintf("%s hours %s minutes %s seconds", hours, minutes, seconds)

	return d
}

// ShortWords function returns duration as a string in format 'Xh Ym Zs'
// Only non-zero values will be taken into account
func ShortWords(duration time.Duration) string {

	var d string
	h, m, s := extractValues(duration)

	seconds := fmt.Sprintf("%d", s)
	d = fmt.Sprintf("%ss", seconds)

	if duration.Minutes() < 1 {
		return d
	}

	minutes := fmt.Sprintf("%d", m)
	d = fmt.Sprintf("%sm %ss", minutes, seconds)

	if duration.Hours() < 1 {
		return d
	}

	hours := fmt.Sprintf("%d", h)
	d = fmt.Sprintf("%sh %sm %ss", hours, minutes, seconds)

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
