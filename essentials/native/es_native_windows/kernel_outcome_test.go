package es_native_windows

import (
	"errors"
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated/eoutcome"
	"testing"
)

func TestNewKernelOutcomeMaySuccess(t *testing.T) {
	oc := NewKernelOutcomeNoObviousError(nil)
	if eoutcome.AssertUnconfirmedOutcomeHasObviousError(oc) {
		t.Error()
	}
	if oc.IsCouldNotResolveProc() {
		t.Error()
	}
}

func TestNewKernelOutcomeCouldNotResolveProc(t *testing.T) {
	e := errors.New("could not resolve proc")
	oc := NewKernelOutcomeCouldNotResolveProc(e)
	if eoutcome.AssertUnconfirmedOutcomeNoObviousError(oc) {
		t.Error()
	}
	if !oc.IsCouldNotResolveProc() {
		t.Error()
	}
	if oc.Cause() != e {
		t.Error()
	}
}
