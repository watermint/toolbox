package ut_compare

import "time"

func Clone(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func ClonePtr(t *time.Time) *time.Time {
	if t == nil {
		return nil
	} else {
		c := Clone(*t)
		return &c
	}
}

// Find earliest time form the list. Returns zero instance if no data.
func Earliest(times ...time.Time) (earliest time.Time) {
	switch len(times) {
	case 0:
		return
	case 1:
		return times[0]
	default:
		earliest = times[0]
		for i := 1; i < len(times); i++ {
			if earliest.After(times[i]) {
				earliest = times[i]
			}
		}
		return
	}
}

// Same as Earliest, but accepts null.
func EarliestPtr(times ...*time.Time) (earliest *time.Time) {
	times1 := make([]time.Time, 0)
	for i := range times {
		if times[i] != nil {
			times1 = append(times1, Clone(*times[i]))
		}
	}
	if e := Earliest(times1...); e.IsZero() {
		return nil
	} else {
		return &e
	}
}

func Latest(times ...time.Time) (latest time.Time) {
	switch len(times) {
	case 0:
		return
	case 1:
		return times[0]
	default:
		latest = times[0]
		for i := 1; i < len(times); i++ {
			if latest.Before(times[i]) {
				latest = times[i]
			}
		}
		return
	}
}

// Same as Latest, but accepts nil.
func LatestPtr(times ...*time.Time) (latest *time.Time) {
	times1 := make([]time.Time, 0)
	for _, t := range times {
		if t != nil {
			times1 = append(times1, *t)
		}
	}
	if e := Latest(times1...); e.IsZero() {
		return nil
	} else {
		return &e
	}
}
