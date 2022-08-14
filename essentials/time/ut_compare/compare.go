package ut_compare

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"time"
)

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

// IsBetween true if `t` is between `a` and `b` (inclusive).
func IsBetween(t, a, b time.Time) bool {
	if b.Before(a) {
		return IsBetween(t, b, a)
	}
	if a.Equal(t) || b.Equal(t) {
		return true
	}
	return a.Before(t) && b.After(t)
}

// IsBetweenOptional true if `t` is (1) between `a` and `b`, (2) `a` < `t` and `b` is zero,
// (3) `t` < `b` and `a` is zero, and (4) both `a` and `b` is zero.
// Always returns false if `b` < `a` and both `a`/`b` is non zero.
func IsBetweenOptional(t time.Time, a, b mo_time.TimeOptional) bool {
	switch {
	case a.IsZero() && b.IsZero():
		return true

	case !a.IsZero() && b.IsZero(): // `a` < `t`, `b` is zero
		at := a.Time()
		if at.Equal(t) {
			return true
		}
		return at.Before(t)

	case a.IsZero() && !b.IsZero(): // `t` < `b`, `a` is zero
		bt := b.Time()
		if bt.Equal(t) {
			return true
		}
		return bt.After(t)

	default:
		// always return false in case `b` < `a`.
		if a.Time().After(b.Time()) {
			return false
		}
		at := a.Time()
		bt := b.Time()

		if at.Equal(t) || bt.Equal(t) {
			return true
		}
		return at.Before(t) && bt.After(t)
	}
}
