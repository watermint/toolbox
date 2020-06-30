package es_timeout

import (
	"context"
	"testing"
	"time"
)

func TestDoWithTimeout(t *testing.T) {
	// Run immediately
	{
		v := 99
		DoWithTimeout(10*time.Millisecond, func(ctx context.Context) {
			v = 1
		})
		if v != 1 {
			t.Error(v)
		}
	}

	// Timeout
	{
		v := 99
		DoWithTimeout(10*time.Millisecond, func(ctx context.Context) {
			time.Sleep(1 * time.Second)
			v = 1
		})
		if v != 99 {
			t.Error(v)
		}
	}
}
