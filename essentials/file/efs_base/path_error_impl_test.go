package efs_base

import (
	"errors"
	"testing"
	"time"
)

func TestPathErrorImpl_IsPathInvalidName(t *testing.T) {
	cause := errors.New(time.Now().String())
	pe := NewPathError(PathErrorInvalidName, cause)
	if !pe.IsPathInvalidName() {
		t.Error()
	}
	if pe.IsPathTooLong() {
		t.Error()
	}
}

func TestPathErrorImpl_IsPathTooLong(t *testing.T) {
	cause := errors.New(time.Now().String())
	pe := NewPathError(PathErrorTooLong, cause)
	if pe.IsPathInvalidName() {
		t.Error()
	}
	if !pe.IsPathTooLong() {
		t.Error()
	}
}
