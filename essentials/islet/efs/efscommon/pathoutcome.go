package efscommon

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/essentials/islet/efs"
	"github.com/watermint/toolbox/essentials/islet/eidiom/eoutcome"
)

const (
	childOutcomeSuccess = iota
	childOutcomeNameOutcome
	childOutcomePathTooLong
)

func NewChildOutcomeSuccess() efs.ChildOutcome {
	return &childOutcomeImpl{
		OutcomeBase: eoutcome.OutcomeBase{Err: nil},
		reason:      childOutcomeSuccess,
	}
}

func NewChildOutcomeByNameOutcome(noc efs.NameOutcome) efs.ChildOutcome {
	return &childOutcomeImpl{
		OutcomeBase: eoutcome.OutcomeBase{Err: noc.Cause()},
		nameOutcome: noc,
		reason:      childOutcomeNameOutcome,
	}
}

func NewChildOutcomePathTooLong(given, allowed int) efs.ChildOutcome {
	reason := errors.New(fmt.Sprintf("path too long: given length (%d), allowed length (%d)", given, allowed))
	return &childOutcomeImpl{
		OutcomeBase: eoutcome.OutcomeBase{Err: reason},
		reason:      childOutcomePathTooLong,
	}
}

type childOutcomeImpl struct {
	eoutcome.OutcomeBase
	nameOutcome efs.NameOutcome
	reason      int
}

func (z childOutcomeImpl) IsInvalidChar() bool {
	return z.nameOutcome != nil && z.nameOutcome.IsInvalidChar()
}

func (z childOutcomeImpl) IsNameReserved() bool {
	return z.nameOutcome != nil && z.nameOutcome.IsNameReserved()
}

func (z childOutcomeImpl) IsNameTooLong() bool {
	return z.nameOutcome != nil && z.nameOutcome.IsNameTooLong()
}

func (z childOutcomeImpl) IsPathTooLong() bool {
	return z.reason == childOutcomePathTooLong
}
