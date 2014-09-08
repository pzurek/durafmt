package durafmt

import (
	"testing"
	"time"
)

var ins = []time.Duration{
	time.Duration(time.Second * 12),
	time.Duration(time.Minute*36 + time.Second*45),
	time.Duration(time.Hour*42 + time.Minute*42 + time.Second*42),
}

func TestHMS(t *testing.T) {

	outs := []string{
		"00:00:12",
		"00:36:45",
		"42:42:42",
	}

	for i, in := range ins {
		d := HMS(in)
		if d != outs[i] {
			t.Errorf("got: %s, expected %s", d, outs[i])
		}
	}
}

func TestLongWords(t *testing.T) {
	outs := []string{
		"12 seconds",
		"36 minutes 45 seconds",
		"42 hours 42 minutes 42 seconds",
	}

	for i, in := range ins {
		d := LongWords(in)
		if d != outs[i] {
			t.Errorf("got: %s, expected %s", d, outs[i])
		}
	}
}

func TestShortWords(t *testing.T) {
	outs := []string{
		"12s",
		"36m 45s",
		"42h 42m 42s",
	}

	for i, in := range ins {
		d := ShortWords(in)
		if d != outs[i] {
			t.Errorf("got: %s, expected %s", d, outs[i])
		}
	}
}
