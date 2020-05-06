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
	{
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
	}

	{
		var buf bytes.Buffer
		l1 := New(LevelDebug, FlavorFileStandard, &buf)
		tee := NewTee()
		tee.AddSubscriber(l1)

		err := EnsureCallerSkip(l1, "msg", "caller", func() string {
			return buf.String()
		})
		if err != nil {
			t.Error(err)
		}
	}

	{
		var buf1 bytes.Buffer
		l1 := New(LevelDebug, FlavorFileStandard, &buf1)
		tee := NewTee()
		tee.AddSubscriber(l1)

		var buf2 bytes.Buffer
		l2 := NewLogCloser(LevelDebug, FlavorFileStandard, es_close.NewNopWriteCloser(&buf2))
		tee.AddSubscriber(l2)
		err := EnsureCallerSkip(l2, "msg", "caller", func() string {
			return buf2.String()
		})
		if err != nil {
			t.Error(err)
		}
		err = EnsureCallerSkip(tee, "msg", "caller", func() string {
			return buf1.String()
		})
		if err != nil {
			t.Error(err)
		}
	}

	{
		var buf1 bytes.Buffer
		l1 := New(LevelDebug, FlavorFileStandard, &buf1)
		tee := NewTee()
		tee.AddSubscriber(l1)

		err := EnsureCallerSkip(tee, "msg", "caller", func() string {
			return buf1.String()
		})
		if err != nil {
			t.Error(err)
		}

		t1 := tee.With(String("t1", "T1"))

		err = EnsureCallerSkip(t1, "msg", "caller", func() string {
			return buf1.String()
		})
		if err != nil {
			t.Error(err)
		}

		var buf2 bytes.Buffer
		l2 := NewLogCloser(LevelDebug, FlavorFileStandard, es_close.NewNopWriteCloser(&buf2))
		tee.AddSubscriber(l2)

		t2 := tee.With(String("hello", "world"))

		err = EnsureCallerSkip(t2, "msg", "caller", func() string {
			return buf2.String()
		})
		if err != nil {
			t.Error(err)
		}

		var buf3 bytes.Buffer
		l3 := NewLogCloser(LevelDebug, FlavorFileStandard, es_close.NewNopWriteCloser(&buf3))
		tee.AddSubscriber(l3)

		t3 := tee.With(String("hello", "world"))

		err = EnsureCallerSkip(t3, "msg", "caller", func() string {
			return buf3.String()
		})
		if err != nil {
			t.Error(err)
		}
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

func TestCallerLogCloser2(t *testing.T) {
	var buf bytes.Buffer

	l := NewLogCloser(LevelDebug, FlavorFileStandard, es_close.NewNopWriteCloser(&buf))
	ll := l.With(String("key", "value"))
	err := EnsureCallerSkip(ll, "msg", "caller", func() string {
		return buf.String()
	})
	if err != nil {
		t.Error(err)
	}
}
