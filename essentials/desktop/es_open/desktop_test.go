package es_open

import (
	"os"
	"testing"
)

func TestCurrentDesktop(t *testing.T) {
	if !testing.Verbose() {
		t.Skip("Skip test")
	}
	d := CurrentDesktop()
	p, err := os.MkdirTemp("", "desktop")
	if err != nil {
		t.Error(err)
		return
	}
	err = d.Open(p)
	switch {
	case err == nil:
		t.Log("success")
	case IsOpenFailure(err):
		t.Log("open failure", err)
	case IsOperationUnsupported(err):
		t.Log("unsupported", err)
	}
}
