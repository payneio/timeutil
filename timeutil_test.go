package timeutil

import (
	"testing"
	"time"
)

func TestBeginningOfDay(t *testing.T) {
	in := time.Date(2014, 9, 11, 1, 23, 04, 0, time.UTC)
	out := time.Date(2014, 9, 11, 0, 0, 0, 0, time.UTC)
	if x := BeginningOfDay(in); x != out {
		t.Errorf("BeginningOfDay(%v) = %v, want %v", in, x, out)
	}
}

func TestEndOfDay(t *testing.T) {
	in := time.Date(2014, 9, 11, 1, 23, 04, 0, time.UTC)
	out := time.Date(2014, 9, 11, 23, 59, 59, 999999999, time.UTC)
	if x := EndOfDay(in); x != out {
		t.Errorf("EndOfDay(%v) = %v, want %v", in, x, out)
	}
}

func TestDurationToEndOfDayFrom(t *testing.T) {
	in := time.Date(2014, 9, 11, 23, 59, 59, 111111111, time.UTC)
	out := time.Duration(888888889 * time.Nanosecond)
	if x := DurationToEndOfDayFrom(in); x != out {
		t.Errorf("DurationToEndOfDayFrom(%v) = %v, want %v", in, x, out)
	}

	in = time.Date(2014, 9, 11, 23, 59, 59, 0, time.UTC)
	out = time.Duration(1 * time.Second)
	if x := DurationToEndOfDayFrom(in); x != out {
		t.Errorf("DurationToEndOfDayFrom(%v) = %v, want %v", in, x, out)
	}
}
