package es_stdout

import (
	"testing"
)

func TestNewDefaultOut(t *testing.T) {
	{
		w := newDefaultOut(true, false)
		w.Write([]byte("hello"))
		w.Close()
	}
	{
		w := newDefaultOut(false, false)
		w.Write([]byte("Hello"))
		w.Close()
	}
}
