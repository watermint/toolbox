package es_mutex

import (
	"errors"
	"testing"
	"time"
)

func TestMutexImpl_Do(t *testing.T) {
	m := New()
	m.Do(func() {})
	var err error
	m.Do(func() {
		err = errors.New("some error")
	})
	if err == nil {
		t.Error(err)
	}
}

func TestMutexImpl_Do2(t *testing.T) {
	v := 0
	m := NewWithTimeoutRetry(10*time.Millisecond, 1, func() {
		v = 4
	})
	m.Do(func() {
		v = 1
		ts := time.Now()
		m.Do(func() {
			// should not dead lock
			v = 2
		})
		te := time.Now()
		if te.Sub(ts).Milliseconds() < 10 {
			t.Error(te, ts)
		}
	})
	if v != 4 {
		t.Error(v)
	}
}
