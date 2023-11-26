package eoutcome

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated"
)

type UnconfirmedOutcomeBase struct {
	ObviousErr error
}

func (z UnconfirmedOutcomeBase) HasError() bool {
	return z.ObviousErr != nil
}

func (z UnconfirmedOutcomeBase) IfError(f func() es_idiom_deprecated.UnconfirmedOutcome) es_idiom_deprecated.UnconfirmedOutcome {
	if z.ObviousErr != nil {
		return f()
	}
	return z
}

func (z UnconfirmedOutcomeBase) Cause() error {
	return z.ObviousErr
}

// AssertUnconfirmedOutcomeNoObviousError Returns false if the outcome does not comply with behaviour for no obvious error
func AssertUnconfirmedOutcomeNoObviousError(outcome es_idiom_deprecated.UnconfirmedOutcome) bool {
	if outcome.HasError() {
		return false
	}
	if outcome.Cause() != nil {
		return false
	}
	p := false
	outcome.IfError(func() es_idiom_deprecated.UnconfirmedOutcome {
		p = true
		return outcome
	})
	if p {
		return false
	}
	return true
}

// AssertUnconfirmedOutcomeHasObviousError Returns false if the outcome does not comply with behaviour for obvious error
func AssertUnconfirmedOutcomeHasObviousError(outcome es_idiom_deprecated.UnconfirmedOutcome) bool {
	if !outcome.HasError() {
		return false
	}
	if outcome.Cause() == nil {
		return false
	}
	p := false
	outcome.IfError(func() es_idiom_deprecated.UnconfirmedOutcome {
		p = true
		return outcome
	})
	if !p {
		return false
	}
	return true
}
