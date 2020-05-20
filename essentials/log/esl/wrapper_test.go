package esl

import "testing"

func TestZapWrapper(t *testing.T) {
	l := newDefault()
	l.Debug("test debug", String("key", "val"))
	l.Info("test info", Int("key", 123))
	l.Warn("test warn", Any("key", 456))
	l.Error("test err", Uint("key", 789))

	ll := l.With(Bool("with", true))
	ll.Debug("test debug", String("key", "val"))
	ll.Info("test info", Int("key", 123))
	ll.Warn("test warn", Any("key", 456))
	ll.Error("test err", Uint("key", 789))

	if err := l.Sync(); err != nil {
		t.Error(err)
	}
}
