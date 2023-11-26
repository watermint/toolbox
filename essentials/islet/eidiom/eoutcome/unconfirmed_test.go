package eoutcome

import (
	"errors"
	"github.com/watermint/toolbox/essentials/islet/eidiom"
	"testing"
)

func TestUnconfirmedOutcomeBase_Cause(t *testing.T) {
	{
		uoc := UnconfirmedOutcomeBase{ObviousErr: nil}
		if x := uoc.Cause(); x != nil {
			t.Error(x)
		}
	}
	{
		e := errors.New("unconfirmed")
		uoc := UnconfirmedOutcomeBase{ObviousErr: e}
		if x := uoc.Cause(); x != e {
			t.Error(x)
		}
	}
}

func TestUnconfirmedOutcomeBase_HasError(t *testing.T) {
	{
		uoc := UnconfirmedOutcomeBase{ObviousErr: nil}
		if x := uoc.HasError(); x {
			t.Error(x)
		}
	}
	{
		e := errors.New("unconfirmed")
		uoc := UnconfirmedOutcomeBase{ObviousErr: e}
		if x := uoc.HasError(); !x {
			t.Error(x)
		}
	}
}

func TestUnconfirmedOutcomeBase_IfError(t *testing.T) {
	{
		uoc := UnconfirmedOutcomeBase{ObviousErr: nil}
		p := false
		uoc.IfError(func() eidiom.UnconfirmedOutcome {
			p = true
			return uoc
		})
		if p {
			t.Error(p)
		}
	}
	{
		e := errors.New("unconfirmed")
		uoc := UnconfirmedOutcomeBase{ObviousErr: e}
		p := false
		uoc.IfError(func() eidiom.UnconfirmedOutcome {
			p = true
			return uoc
		})
		if !p {
			t.Error(p)
		}
	}
}

func TestAssertUnconfirmedOutcomeNoObviousError(t *testing.T) {
	{
		uoc := UnconfirmedOutcomeBase{ObviousErr: nil}
		if x := AssertUnconfirmedOutcomeNoObviousError(uoc); !x {
			t.Error(x)
		}
	}
	{
		e := errors.New("unconfirmed")
		uoc := UnconfirmedOutcomeBase{ObviousErr: e}
		if x := AssertUnconfirmedOutcomeNoObviousError(uoc); x {
			t.Error(x)
		}
	}
}

func TestAssertUnconfirmedOutcomeHasObviousError(t *testing.T) {
	{
		uoc := UnconfirmedOutcomeBase{ObviousErr: nil}
		if x := AssertUnconfirmedOutcomeHasObviousError(uoc); x {
			t.Error(x)
		}
	}
	{
		e := errors.New("unconfirmed")
		uoc := UnconfirmedOutcomeBase{ObviousErr: e}
		if x := AssertUnconfirmedOutcomeHasObviousError(uoc); !x {
			t.Error(x)
		}
	}
}
