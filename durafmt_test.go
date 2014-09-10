package durafmt

import (
	"testing"
	"time"
)

var ins = []time.Duration{
	time.Duration(time.Second * 1),
	time.Duration(time.Second * 12),
	time.Duration(time.Minute * 25),
	time.Duration(time.Minute*1 + time.Second*12),
	time.Duration(time.Minute*36 + time.Second*45),
	time.Duration(time.Hour*67 + time.Second*25),
	time.Duration(time.Hour*1 + time.Minute*1 + time.Second*1),
	time.Duration(time.Hour*42 + time.Minute*42 + time.Second*42),
}

func TestHM(t *testing.T) {

	outs := []string{
		"00:00",
		"00:00",
		"00:25",
		"00:01",
		"00:36",
		"67:00",
		"01:01",
		"42:42",
	}

	for i, in := range ins {
		d := HM(in)
		if d != outs[i] {
			t.Errorf("got: %s, expected %s", d, outs[i])
		}
	}
}

func TestHMWithSeparator(t *testing.T) {

	outs := []string{
		"00.00",
		"00.00",
		"00.25",
		"00.01",
		"00.36",
		"67.00",
		"01.01",
		"42.42",
	}

	for i, in := range ins {
		d := HMWithSeparator(in, ".")
		if d != outs[i] {
			t.Errorf("got: %s, expected %s", d, outs[i])
		}
	}
}

func TestHMS(t *testing.T) {

	outs := []string{
		"00:00:01",
		"00:00:12",
		"00:25:00",
		"00:01:12",
		"00:36:45",
		"67:00:25",
		"01:01:01",
		"42:42:42",
	}

	for i, in := range ins {
		d := HMS(in)
		if d != outs[i] {
			t.Errorf("got: %s, expected %s", d, outs[i])
		}
	}
}

func TestHMSWithSeparator(t *testing.T) {

	outs := []string{
		"00.00.01",
		"00.00.12",
		"00.25.00",
		"00.01.12",
		"00.36.45",
		"67.00.25",
		"01.01.01",
		"42.42.42",
	}

	for i, in := range ins {
		d := HMSWithSeparator(in, ".")
		if d != outs[i] {
			t.Errorf("got: %s, expected %s", d, outs[i])
		}
	}
}

func TestLongWords(t *testing.T) {
	outs := []string{
		"1 second",
		"12 seconds",
		"25 minutes 0 seconds",
		"1 minute 12 seconds",
		"36 minutes 45 seconds",
		"67 hours 0 minutes 25 seconds",
		"1 hour 1 minute 1 second",
		"42 hours 42 minutes 42 seconds",
	}

	for i, in := range ins {
		d := LongWords(in)
		if d != outs[i] {
			t.Errorf("got: %s, expected %s", d, outs[i])
		}
	}
}

func TestLongWordsWithSeparator(t *testing.T) {
	outs := []string{
		"1 second",
		"12 seconds",
		"25 minutes, 0 seconds",
		"1 minute, 12 seconds",
		"36 minutes, 45 seconds",
		"67 hours, 0 minutes, 25 seconds",
		"1 hour, 1 minute, 1 second",
		"42 hours, 42 minutes, 42 seconds",
	}

	for i, in := range ins {
		d := LongWordsWithSeparator(in, ",")
		if d != outs[i] {
			t.Errorf("got: %s, expected %s", d, outs[i])
		}
	}
}

func TestShortWords(t *testing.T) {
	outs := []string{
		"1s",
		"12s",
		"25m 0s",
		"1m 12s",
		"36m 45s",
		"67h 0m 25s",
		"1h 1m 1s",
		"42h 42m 42s",
	}

	for i, in := range ins {
		d := ShortWords(in)
		if d != outs[i] {
			t.Errorf("got: %s, expected %s", d, outs[i])
		}
	}
}
