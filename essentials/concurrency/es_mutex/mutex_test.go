package es_mutex

import (
	"errors"
	"testing"
)

func TestMutexImpl_Do(t *testing.T) {
	m := New()
	m.Do(func() {})
	var err error
	m.Do(func() {
		err = errors.New("some error")
	})
	if err == nil {
		t.Error(err)
	}
}
