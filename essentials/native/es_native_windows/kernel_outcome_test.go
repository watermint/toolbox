package es_native_windows

import (
	"errors"
	"testing"
)

func TestNewKernelOutcomeMaySuccess(t *testing.T) {
	err := NewKernelOutcomeNoObviousError(nil)
	if err != nil {
		t.Error("Expected nil error, got:", err)
	}
}

func TestNewKernelOutcomeCouldNotResolveProc(t *testing.T) {
	origErr := errors.New("could not resolve proc")
	err := NewKernelOutcomeCouldNotResolveProc(origErr)

	if err == nil {
		t.Error("Expected non-nil error")
	}

	if !IsCouldNotResolveProcError(err) {
		t.Error("Expected to be CouldNotResolveProc error")
	}

	if err.Error() != origErr.Error() {
		t.Error("Error message mismatch")
	}
}
