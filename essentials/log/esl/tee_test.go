package esl

import (
	"github.com/watermint/toolbox/essentials/io/es_close"
	"testing"
)

func TestNewTee(t *testing.T) {
	buf1 := es_close.NewNopCloseBuffer()
	buf2 := es_close.NewNopCloseBuffer()
	l1 := New(LevelDebug, FlavorFileStandard, buf1)
	l2 := New(LevelDebug, FlavorFileStandard, buf2)
	tee := NewTee()

	tw := func(l Logger) {
		l.Debug("test debug", String("key", "val"))
		l.Info("test info", Int("key", 123))
		l.Warn("test warn", Any("key", 456))
		l.Error("test err", Uint("key", 789))
		if err := l.Sync(); err != nil {
			t.Error(err)
		}
	}

	// should not write to anywhere
	tw(tee)
	if buf1.String() != "" || buf2.String() != "" {
		t.Error(buf1, buf2)
	}

	tee.AddSubscriber(l1)
	tw(tee)
	// write only on l1
	if buf1.String() == "" || buf2.String() != "" {
		t.Error(buf1, buf2)
	}

	tee.AddSubscriber(l2)
	tw(tee)
	// write on both l1 and l2
	if buf1.String() == "" || buf2.String() == "" {
		t.Error(buf1, buf2)
	}
}
