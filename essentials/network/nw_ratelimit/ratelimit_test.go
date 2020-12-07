package nw_ratelimit

import (
	"errors"
	"testing"
	"time"
)

func TestLimitStateImpl_AddError(t *testing.T) {
	shortWait := 100 * time.Millisecond
	longWait := 200 * time.Millisecond
	newState := func() *limitStateImpl {
		ls := newLimitState().(*limitStateImpl)
		ls.durationShortWait = shortWait
		ls.durationLongWait = longWait
		ls.maxLastErrors = 5
		ls.sameErrorThreshold = 4
		return ls
	}

	add := func(ls LimitState, err error) (abort bool, latency time.Duration) {
		ts := time.Now()
		abort = ls.AddError("123", "end/point", err)
		te := time.Now()
		return abort, te.Sub(ts)
	}

	// Single add: should short waited
	{
		ls := newState()
		a, l := add(ls, errors.New("add-error"))
		if a {
			t.Error("invalid abort")
		}
		if l < shortWait || longWait < l {
			t.Error("invalid wait range")
		}
	}

	// Promote retryAction to longWait
	{
		ls := newState()
		for i := 0; i <= ls.sameErrorThreshold; i++ {
			a, l := add(ls, errors.New("add-error"))
			if a {
				t.Error("invalid abort")
			}
			if l < shortWait || longWait < l {
				t.Error("invalid wait range")
			}
		}

		// should promote
		a, l := add(ls, errors.New("add-error"))
		if a {
			t.Error("invalid abort")
		}
		if l < longWait {
			t.Error("invalid wait range")
		}
	}

	// Promote retryAction to longWait
	{
		ls := newState()
		for i := 0; i <= ls.sameErrorThreshold*2; i++ {
			a, l := add(ls, errors.New("add-error"))
			if a {
				t.Error("invalid abort")
			}
			if i <= ls.sameErrorThreshold && (l < shortWait || longWait < l) {
				t.Error("invalid wait range")
			}
			if ls.sameErrorThreshold < i && l < longWait {
				t.Error("invalid wait range")
			}
		}

		// should promote
		a, _ := add(ls, errors.New("add-error"))
		if !a {
			t.Error("should abort")
		}
	}
}

func TestLimitStateImpl_UpdateRetryAfter(t *testing.T) {
	wait := func(ls LimitState) time.Duration {
		ts := time.Now()
		ls.WaitIfRequired("123", "end/point")
		te := time.Now()
		return te.Sub(ts)
	}

	// No retry after
	{
		ls := newLimitState()
		d := wait(ls)
		if 1*time.Millisecond < d {
			t.Error("invalid wait")
		}
	}

	// should not effect to other hash
	{
		ls := newLimitState()
		ls.UpdateRetryAfter("456", "end/point", time.Now().Add(10*time.Millisecond))
		d := wait(ls)
		if 1*time.Millisecond < d {
			t.Error("invalid wait")
		}
	}

	// should wait for retry after
	{
		ls := newLimitState()
		ls.UpdateRetryAfter("123", "end/point", time.Now().Add(10*time.Millisecond))
		d := wait(ls)
		if d < 10*time.Millisecond {
			t.Error("invalid wait")
		}
	}
}
