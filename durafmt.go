package durafmt

import (
	"fmt"
	"time"
)

// HMWithSeparator function returns duration as a string in format 'hh:mm'.
func HMWithSeparator(duration time.Duration, sep string) string {

	h, m, _ := extractValues(duration)

	hours := fmt.Sprintf("%0.2d", h)
	minutes := fmt.Sprintf("%0.2d", m)

	d := fmt.Sprintf("%s%s%s", hours, sep, minutes)

	return d
}

// HM function returns duration as a string in format 'hh:mm'.
func HM(duration time.Duration) string {
	return HMWithSeparator(duration, ":")
}

// HMSWithSeparator function returns duration as a string in format 'hh:mm:ss'.
func HMSWithSeparator(duration time.Duration, sep string) string {

	h, m, s := extractValues(duration)

	hours := fmt.Sprintf("%0.2d", h)
	minutes := fmt.Sprintf("%0.2d", m)
	seconds := fmt.Sprintf("%0.2d", s)

	d := fmt.Sprintf("%s%s%s%s%s", hours, sep, minutes, sep, seconds)

	return d
}

// HMS function returns duration as a string in format 'hh:mm:ss'.
func HMS(duration time.Duration) string {
	return HMSWithSeparator(duration, ":")
}

// LongWordsWithSeparator function returns duration as a string in format 'X hours Y minutes Z seconds'.
//
// Only non-zero values will be taken into account.
func LongWordsWithSeparator(duration time.Duration, sep string) string {

	var d string
	h, m, s := extractValues(duration)

	seconds := fmt.Sprintf("%d", s)
	d = fmt.Sprintf("%s seconds", seconds)

	if duration.Minutes() < 1 {
		return d
	}

	minutes := fmt.Sprintf("%d", m)
	d = fmt.Sprintf("%s minutes%s %s seconds", minutes, sep, seconds)

	if duration.Hours() < 1 {
		return d
	}

	hours := fmt.Sprintf("%d", h)
	d = fmt.Sprintf("%s hours%s %s minutes%s %s seconds", hours, sep, minutes, sep, seconds)

	return d
}

// LongWords function returns duration as a string in format 'X hours Y minutes Z seconds'.
//
// Only non-zero values will be taken into account.
func LongWords(duration time.Duration) string {
	return LongWordsWithSeparator(duration, "")
}

// ShortWords function returns duration as a string in format 'Xh Ym Zs'.
//
// Only non-zero values will be taken into account.
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
