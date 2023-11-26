package eoutcome

import "testing"

func TestNewParseSuccess(t *testing.T) {
	oc := NewParseSuccess()
	AssertOutcomeSuccess(oc)
	if oc.IsInvalidChar() {
		t.Error()
	}
	if oc.IsInvalidFormat() {
		t.Error()
	}
}

func TestNewParseInvalidChar(t *testing.T) {
	oc := NewParseInvalidChar("must use allowed chars")
	AssertOutcomeFailure(oc)
	if !oc.IsInvalidChar() {
		t.Error()
	}
	if oc.IsInvalidFormat() {
		t.Error()
	}
}

func TestNewParseInvalidFormat(t *testing.T) {
	oc := NewParseInvalidFormat("must comply certain format")
	AssertOutcomeFailure(oc)
	if oc.IsInvalidChar() {
		t.Error()
	}
	if !oc.IsInvalidFormat() {
		t.Error()
	}
}
