package esl

import (
	"github.com/watermint/toolbox/essentials/io/es_close"
	"testing"
)

func TestNewLogCloser(t *testing.T) {
	buf := es_close.NewNopCloseBuffer()
	l := NewLogCloser(LevelDebug, FlavorFileStandard, buf)
	l.Debug("test debug", String("key", "val"))
	l.Info("test info", Int("key", 123))
	l.Warn("test warn", Any("key", 456))
	l.Error("test err", Uint("key", 789))
	if err := l.Sync(); err != nil {
		t.Error(err)
	}
	if err := l.Close(); err != nil {
		t.Error(err)
	}
}
