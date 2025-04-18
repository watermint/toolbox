package es_timeout

import (
	"io"
	"testing"
	"time"
)

type waitWriter struct {
}

func (w waitWriter) Write(p []byte) (n int, err error) {
	time.Sleep(1 * time.Second)
	return len(p), nil
}

func TestToWriter_Write(t *testing.T) {
	// Write return immediately
	{
		tw := New(io.Discard, 10*time.Millisecond)
		n, _ := tw.Write([]byte("hello"))
		if n != 5 {
			t.Error(n)
		}
	}

	// Write returns after 1 second
	{
		tw := New(&waitWriter{}, 10*time.Millisecond)
		n, err := tw.Write([]byte("hello"))
		if n != 0 || err != ErrorTimeout {
			t.Error(n, err)
		}
	}
}
