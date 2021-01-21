package es_file_random

import (
	"io"
	"testing"
)

func TestNewReader(t *testing.T) {
	{
		r := NewReader(128)
		buf := make([]byte, 64)

		if n, err := r.Read(buf); n != 64 || err != nil {
			t.Error(n, err)
		}
		if n, err := r.Read(buf); n != 64 || err != io.EOF {
			t.Error(n, err)
		}
	}

	{
		r := NewReader(96)
		buf := make([]byte, 64)

		if n, err := r.Read(buf); n != 64 || err != nil {
			t.Error(n, err)
		}
		if n, err := r.Read(buf); n != 32 || err != io.EOF {
			t.Error(n, err)
		}
	}
}
