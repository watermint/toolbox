package efscommon

import (
	"errors"
	"fmt"
	efs_deprecated2 "github.com/watermint/toolbox/essentials/file/efs_alpha"
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated/eoutcome"
)

const (
	childOutcomeSuccess = iota
	childOutcomeNameOutcome
	childOutcomePathTooLong
)

func NewChildOutcomeSuccess() efs_deprecated2.ChildOutcome {
	return &childOutcomeImpl{
		OutcomeBase: eoutcome.OutcomeBase{Err: nil},
		reason:      childOutcomeSuccess,
	}
}

func NewChildOutcomeByNameOutcome(noc efs_deprecated2.NameOutcome) efs_deprecated2.ChildOutcome {
	return &childOutcomeImpl{
		OutcomeBase: eoutcome.OutcomeBase{Err: noc.Cause()},
		nameOutcome: noc,
		reason:      childOutcomeNameOutcome,
	}
}

func NewChildOutcomePathTooLong(given, allowed int) efs_deprecated2.ChildOutcome {
	reason := errors.New(fmt.Sprintf("path too long: given length (%d), allowed length (%d)", given, allowed))
	return &childOutcomeImpl{
		OutcomeBase: eoutcome.OutcomeBase{Err: reason},
		reason:      childOutcomePathTooLong,
	}
}

type childOutcomeImpl struct {
	eoutcome.OutcomeBase
	nameOutcome efs_deprecated2.NameOutcome
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
