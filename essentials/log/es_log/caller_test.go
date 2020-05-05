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

	l1 := New(LevelDebug, FlavorFileStandard, &buf)
	tee := NewTee()
	tee.AddSubscriber(l1)

	err := EnsureCallerSkip(tee, "msg", "caller", func() string {
		return buf.String()
	})
	if err != nil {
		t.Error(err)
	}
	err = EnsureCallerSkip(l1, "msg", "caller", func() string {
		return buf.String()
	})
	if err != nil {
		t.Error(err)
	}

	l2 := NewLogCloser(LevelDebug, FlavorFileStandard, es_close.NewNopWriteCloser(&buf))
	tee.AddSubscriber(l2)
	err = EnsureCallerSkip(l2, "msg", "caller", func() string {
		return buf.String()
	})
	if err != nil {
		t.Error(err)
	}
	err = EnsureCallerSkip(tee, "msg", "caller", func() string {
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
