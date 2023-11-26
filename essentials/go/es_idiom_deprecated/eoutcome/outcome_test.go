package eoutcome

import (
	"errors"
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated"
	"testing"
)

func TestOutcome_Cause(t *testing.T) {
	{
		e := errors.New("cause")
		o := &OutcomeBase{Err: e}
		if x := o.Cause(); x != e {
			t.Error(x, e)
		}
	}

	{
		o := &OutcomeBase{Err: nil}
		if x := o.Cause(); x != nil {
			t.Error(x)
		}
	}
}

func TestOutcome_IfError(t *testing.T) {
	{
		e := errors.New("cause")
		o := &OutcomeBase{Err: e}
		p := false
		x := o.IfError(func() es_idiom_deprecated.Outcome {
			p = true
			return o
		})
		if !p || x != o {
			t.Error(p, x)
		}
	}

	{
		o := &OutcomeBase{Err: nil}
		p := false
		x := o.IfError(func() es_idiom_deprecated.Outcome {
			p = true
			return o
		})
		if p || x.IsError() {
			t.Error(p, x)
		}
	}
}

func TestOutcome_IfOk(t *testing.T) {
	{
		e := errors.New("cause")
		o := &OutcomeBase{Err: e}
		p := false
		o.IfOk(func() {
			p = true
		})
		if p {
			t.Error(p)
		}
	}

	{
		o := &OutcomeBase{Err: nil}
		p := false
		o.IfOk(func() {
			p = true
		})
		if !p {
			t.Error(p)
		}
	}
}

func TestOutcome_String(t *testing.T) {
	{
		e := errors.New("cause")
		o := &OutcomeBase{Err: e}
		if x := o.String(); x != "cause" {
			t.Error(x)
		}
	}

	{
		o := &OutcomeBase{Err: nil}
		if x := o.String(); x != "" {
			t.Error(x)
		}
	}
}

func TestOutcome_IsOk(t *testing.T) {
	{
		e := errors.New("cause")
		o := &OutcomeBase{Err: e}
		if x := o.IsOk(); x {
			t.Error(x)
		}
	}

	{
		o := &OutcomeBase{Err: nil}
		if x := o.IsOk(); !x {
			t.Error(x)
		}
	}
}

func TestOutcome_IsError(t *testing.T) {
	{
		e := errors.New("cause")
		o := &OutcomeBase{Err: e}
		if x := o.IsError(); !x {
			t.Error(x)
		}
	}

	{
		o := &OutcomeBase{Err: nil}
		if x := o.IsError(); x {
			t.Error(x)
		}
	}
}

func TestAssertOutcomeSuccess(t *testing.T) {
	{
		e := errors.New("cause")
		o := &OutcomeBase{Err: e}
		if x := AssertOutcomeSuccess(o); x {
			t.Error(x)
		}
	}

	{
		o := &OutcomeBase{Err: nil}
		if x := AssertOutcomeSuccess(o); !x {
			t.Error(x)
		}
	}
}

func TestAssertOutcomeFailure(t *testing.T) {
	{
		e := errors.New("cause")
		o := &OutcomeBase{Err: e}
		if x := AssertOutcomeFailure(o); !x {
			t.Error(x)
		}
	}

	{
		o := &OutcomeBase{Err: nil}
		if x := AssertOutcomeFailure(o); x {
			t.Error(x)
		}
	}
}
