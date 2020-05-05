package es_log

import (
	"bytes"
	"github.com/watermint/toolbox/essentials/io/es_close"
	"testing"
)

func TestCallerWrapper(t *testing.T) {
	var buf bytes.Buffer

	l := New(LevelDebug, FlavorFileStandard, &buf)
	err := EnsureCallerSkip(l, "msg", "caller", func() string {
		return buf.String()
	})
	if err != nil {
		t.Error(err)
	}
}

func TestCallerTee(t *testing.T) {
	var buf bytes.Buffer

	l := New(LevelDebug, FlavorFileStandard, &buf)
	l2 := NewTee()
	l2.AddSubscriber(l)

	err := EnsureCallerSkip(l2, "msg", "caller", func() string {
		return buf.String()
	})
	if err != nil {
		t.Error(err)
	}
}

func TestCallerLogCloser(t *testing.T) {
	var buf bytes.Buffer

	l := NewLogCloser(LevelDebug, FlavorFileStandard, es_close.NewNopWriteCloser(&buf))
	err := EnsureCallerSkip(l, "msg", "caller", func() string {
		return buf.String()
	})
	if err != nil {
		t.Error(err)
	}
}
