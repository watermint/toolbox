package app_workspace

import (
	"os"
	"strings"
	"testing"
)

func TestNewJobId(t *testing.T) {
	j1 := NewJobId()
	if j1 == "" {
		t.Error(j1)
	}
	j2 := NewJobId()
	if j1 == j2 {
		t.Error(j1, j2)
	}
}

func TestDefaultAppPath(t *testing.T) {
	p, err := DefaultAppPath()
	if err != nil {
		t.Error(p, err)
	}
}

func TestNewWorkspace(t *testing.T) {
	p, err := os.MkdirTemp("", "ws")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()

	ws, err := NewWorkspace(p, false)
	if err != nil {
		t.Error(err)
	}
	if ws.Home() != p {
		t.Error(ws.Home())
	}
	if !strings.HasPrefix(ws.Log(), p) {
		t.Error(ws.Log())
	}
	if !strings.HasPrefix(ws.Test(), p) {
		t.Error(ws.Test())
	}
	if !strings.HasPrefix(ws.Job(), p) {
		t.Error(ws.Job())
	}
	if !strings.HasPrefix(ws.KVS(), p) {
		t.Error(ws.Job())
	}
	if !strings.HasPrefix(ws.Report(), p) {
		t.Error(ws.Report())
	}
	if !strings.HasPrefix(ws.Secrets(), p) {
		t.Error(ws.Secrets())
	}
	somewhere, err := ws.Descendant("somewhere")
	if err != nil || !strings.HasPrefix(somewhere, p) {
		t.Error(somewhere, err)
	}
}
