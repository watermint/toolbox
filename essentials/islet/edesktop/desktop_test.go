package edesktop

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
	oc := d.Open(p)
	switch {
	case oc.IsOk():
		t.Log("success")
	case oc.IsOpenFailure():
		t.Log("open failure", oc.Cause())
	case oc.IsOperationUnsupported():
		t.Log("unsupported", oc.Cause())
	}
}
