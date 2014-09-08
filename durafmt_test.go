package durafmt

import (
	"testing"
	"time"
)

func TestHMS(t *testing.T) {
	duration := time.Duration(time.Hour*2 + time.Minute*54 + time.Second*45)
	d := HMS(duration)
	s := "02:54:45"
	if d != s {
		t.Errorf("HMS: %s, want %s", d, s)
	}
}

func TestLongWords(t *testing.T) {
	duration := time.Duration(time.Hour*2 + time.Minute*54 + time.Second*45)
	d := LongWords(duration)
	s := "2 hours 54 minutes 45 seconds"
	if d != s {
		t.Errorf("HMS: %s, want %s", d, s)
	}
}

func TestShortWords(t *testing.T) {
	duration := time.Duration(time.Hour*2 + time.Minute*54 + time.Second*45)
	d := ShortWords(duration)
	s := "2h 54m 45s"
	if d != s {
		t.Errorf("HMS: %s, want %s", d, s)
	}
}
