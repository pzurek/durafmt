package durafmt

import (
	"fmt"
	"strings"
	"time"
)

// HMWithSeparator function returns duration as a string in format 'hh:mm' with a user defined separator.
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

// HMSWithSeparator function returns duration as a string in format 'hh:mm:ss' with a user defined separator.
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

// WordsWithSeparator function returns duration as a string in format 'X hours Y minutes Z seconds' with a user defined separator.
//
// Only non-zero values will be taken into account.
func WordsWithSeparator(duration time.Duration, sep string) string {

	var d string
	h, m, s := extractValues(duration)

	seconds := fmt.Sprintf("%d", s)
	d = fmt.Sprintf("%s seconds", seconds)

	if duration.Minutes() < 1 {
		return fixSingular(d)
	}

	minutes := fmt.Sprintf("%d", m)
	d = fmt.Sprintf("%s minutes%s %s seconds", minutes, sep, seconds)

	if duration.Hours() < 1 {
		return fixSingular(d)
	}

	hours := fmt.Sprintf("%d", h)
	d = fmt.Sprintf("%s hours%s %s minutes%s %s seconds", hours, sep, minutes, sep, seconds)
	return fixSingular(d)
}

// Replacing all plural forms with signulars where needed
func fixSingular(d string) string {
	s := " " + d
	s = strings.Replace(s, " 1 seconds", " 1 second", -1)
	s = strings.Replace(s, " 1 minutes", " 1 minute", -1)
	s = strings.Replace(s, " 1 hours", " 1 hour", -1)
	s = strings.TrimSpace(s)
	return s
}

// Words function returns duration as a string in format 'X hours Y minutes Z seconds'.
//
// Only non-zero values will be taken into account.
func Words(duration time.Duration) string {
	return WordsWithSeparator(duration, "")
}

// Short function returns duration as a string in format 'Xh Ym Zs'.
//
// Only non-zero values will be taken into account.
func Short(duration time.Duration) string {

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
