package timeutil

import "time"

func BeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, time.UTC).Add(-1 * time.Nanosecond)
}

func DurationOfBeginningOfDayTo(t time.Time) time.Duration {
	return t.Sub(BeginningOfDay(t))
}

func DurationToEndOfDayFrom(t time.Time) time.Duration {
	return EndOfDay(t).Add(1 * time.Nanosecond).Sub(t)
}

func TodayAtSameTimeAs(t time.Time) time.Time {
	n := time.Now() // treat all times as today
	return time.Date(n.Year(), n.Month(), n.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
}
