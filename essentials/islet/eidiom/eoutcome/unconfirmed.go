package eoutcome

import (
	"github.com/watermint/toolbox/essentials/islet/eidiom"
)

type UnconfirmedOutcomeBase struct {
	ObviousErr error
}

func (z UnconfirmedOutcomeBase) HasError() bool {
	return z.ObviousErr != nil
}

func (z UnconfirmedOutcomeBase) IfError(f func() eidiom.UnconfirmedOutcome) eidiom.UnconfirmedOutcome {
	if z.ObviousErr != nil {
		return f()
	}
	return z
}

func (z UnconfirmedOutcomeBase) Cause() error {
	return z.ObviousErr
}

// AssertUnconfirmedOutcomeNoObviousError Returns false if the outcome does not comply with behaviour for no obvious error
func AssertUnconfirmedOutcomeNoObviousError(outcome eidiom.UnconfirmedOutcome) bool {
	if outcome.HasError() {
		return false
	}
	if outcome.Cause() != nil {
		return false
	}
	p := false
	outcome.IfError(func() eidiom.UnconfirmedOutcome {
		p = true
		return outcome
	})
	if p {
		return false
	}
	return true
}

// AssertUnconfirmedOutcomeHasObviousError Returns false if the outcome does not comply with behaviour for obvious error
func AssertUnconfirmedOutcomeHasObviousError(outcome eidiom.UnconfirmedOutcome) bool {
	if !outcome.HasError() {
		return false
	}
	if outcome.Cause() == nil {
		return false
	}
	p := false
	outcome.IfError(func() eidiom.UnconfirmedOutcome {
		p = true
		return outcome
	})
	if !p {
		return false
	}
	return true
}
