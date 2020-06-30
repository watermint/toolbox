package es_timeout

import (
	"context"
	"time"
)

func DoWithTimeout(timeout time.Duration, f func(ctx context.Context)) bool {
	wait := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go func(ctx context.Context) {
		f(ctx)
		wait <- struct{}{}
	}(ctx)

	select {
	case <-ctx.Done():
		return false
	case <-wait:
		return true
	}
}
