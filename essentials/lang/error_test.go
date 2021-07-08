package lang

import (
	"errors"
	"testing"
)

func TestNewMultiError(t *testing.T) {
	e1 := newMultiError(errors.New("abc"), nil, errors.New("def"))
	if e1.NumErrors() != 2 {
		t.Error(e1.NumErrors())
	}

	e2 := newMultiError(nil, nil, errors.New("def"))
	if e2.NumErrors() != 1 {
		t.Error(e2.NumErrors())
	}

	e3 := newMultiError(nil, nil, nil)
	if e3.NumErrors() != 0 {
		t.Error(e3.NumErrors())
	}
}

func TestNewMultiErrorOrNull(t *testing.T) {
	e1 := NewMultiErrorOrNull(errors.New("abc"), nil, errors.New("def")).(*MultiError)
	if e1.NumErrors() != 2 {
		t.Error(e1.NumErrors())
	}

	e2 := NewMultiErrorOrNull(nil, nil, errors.New("def")).(*MultiError)
	if e2.NumErrors() != 1 {
		t.Error(e2.NumErrors())
	}

	e3 := NewMultiErrorOrNull(nil, nil, nil)
	if e3 != nil {
		t.Error(e3)
	}
}
