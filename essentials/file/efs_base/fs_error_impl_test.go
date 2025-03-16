package efs_base

import (
	"errors"
	"testing"
	"time"
)

func TestFsError_IsConflict(t *testing.T) {
	cause := errors.New(time.Now().String())
	fe := NewFsError(FsErrorReasonConflict, cause)
	if !fe.IsConflict() {
		t.Error()
	}
	if fe.IsTimeout() {
		t.Error()
	}
}

func TestFsError_IsNotAllowed(t *testing.T) {
	cause := errors.New(time.Now().String())
	fe := NewFsError(FsErrorReasonNotAllowed, cause)
	if !fe.IsNotAllowed() {
		t.Error()
	}
	if fe.IsConflict() {
		t.Error()
	}
}

func TestFsError_IsPermission(t *testing.T) {
	cause := errors.New(time.Now().String())
	fe := NewFsError(FsErrorReasonPermission, cause)
	if !fe.IsPermission() {
		t.Error()
	}
	if fe.IsConflict() {
		t.Error()
	}
}

func TestFsError_IsTimeout(t *testing.T) {
	cause := errors.New(time.Now().String())
	fe := NewFsError(FsErrorReasonTimeout, cause)
	if !fe.IsTimeout() {
		t.Error()
	}
	if fe.IsConflict() {
		t.Error()
	}
}
