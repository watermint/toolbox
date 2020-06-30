package es_timeout

import (
	"context"
	"errors"
	"github.com/watermint/toolbox/essentials/concurrency/es_timeout"
	"io"
	"time"
)

var (
	ErrorTimeout = errors.New("i/o timeout")
)

type TimeoutWriter interface {
	io.Writer
}

func New(w io.Writer, timeout time.Duration) TimeoutWriter {
	return &toWriter{
		w:       w,
		timeout: timeout,
	}
}

type toWriter struct {
	w       io.Writer
	timeout time.Duration
}

func (z *toWriter) Write(p []byte) (n int, err error) {
	to := es_timeout.DoWithTimeout(z.timeout, func(ctx context.Context) {
		n, err = z.w.Write(p)
	})
	if !to {
		n = 0
		err = ErrorTimeout
	}
	return
}
