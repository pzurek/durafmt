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
