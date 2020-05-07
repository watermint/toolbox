package es_log

import "testing"

func TestNewDefault(t *testing.T) {
	l := newDefault()
	l.Debug("test debug", String("key", "val"))
	l.Info("test info", Int("key", 123))
	l.Warn("test warn", Any("key", 456))
	l.Error("test err", Uint("key", 789))
	if err := l.Sync(); err != nil {
		t.Error(err)
	}
}
