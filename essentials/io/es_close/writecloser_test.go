package es_close

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestWriteCloser_Write(t *testing.T) {
	f, err := os.CreateTemp("", "wc")
	if err != nil {
		t.Error(err)
	}
	p := f.Name()
	w := New(f)

	if n, err := fmt.Fprint(w, "Hello"); err != nil && n != 5 {
		t.Error(n, err)
	}
	if err := w.Close(); err != nil {
		t.Error(err)
	}

	// should not fail after close
	if n, err := fmt.Fprint(w, "World"); err != nil && n != 5 {
		t.Error(n, err)
	}

	g, err := os.ReadFile(p)
	if err != nil {
		t.Error(err)
	}
	// should written only the first write
	if string(g) != "Hello" {
		t.Error(g)
	}
}

func TestNopWriteCloser_Write(t *testing.T) {
	var buf bytes.Buffer
	nw := NewNopWriteCloser(&buf)

	if n, err := fmt.Fprint(nw, "Hello"); err != nil || n != 5 {
		t.Error(n, err)
	}
	if err := nw.Close(); err != nil {
		t.Error(err)
	}
}

func TestNewNopCloseBuffer(t *testing.T) {
	nw := NewNopCloseBuffer()

	if n, err := fmt.Fprint(nw, "Hello"); err != nil || n != 5 {
		t.Error(n, err)
	}
	if err := nw.Close(); err != nil {
		t.Error(err)
	}
	if nw.String() != "Hello" || nw.Len() != 5 {
		t.Error(nw.String(), nw.Len())
	}
	if bytes.Compare(nw.Bytes(), []byte("Hello")) != 0 {
		t.Error(nw.Bytes())
	}
}
