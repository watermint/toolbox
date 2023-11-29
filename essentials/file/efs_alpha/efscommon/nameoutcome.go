package efscommon

import (
	"github.com/watermint/toolbox/essentials/file/efs_alpha"
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated/eoutcome"
)

const (
	nameOutcomeSuccess = iota
	nameOutcomeInvalidChar
	nameOutcomeNameReserved
	nameOutcomeNameTooLong
)

func NewNameOutcomeSuccess() efs_alpha.NameOutcome {
	return &nameOutcome{
		OutcomeBase: eoutcome.OutcomeBase{Err: nil},
		reason:      nameOutcomeSuccess,
	}
}

func NewNameOutcomeInvalidChar(invalidChar string) efs_alpha.NameOutcome {
	return &nameOutcome{
		OutcomeBase: eoutcome.NewOutcomeBaseWithErrMessage("invalid char '%s' found", invalidChar),
		reason:      nameOutcomeInvalidChar,
	}
}

func NewNameOutcomeNameReserved(reserved string) efs_alpha.NameOutcome {
	return &nameOutcome{
		OutcomeBase: eoutcome.NewOutcomeBaseWithErrMessage("reserved keyword '%s' found", reserved),
		reason:      nameOutcomeNameReserved,
	}
}

func NewNameOutcomeNameTooLong(length, max int) efs_alpha.NameOutcome {
	return &nameOutcome{
		OutcomeBase: eoutcome.NewOutcomeBaseWithErrMessage("name is too long (%d). maximum length is %d", length, max),
		reason:      childOutcomePathTooLong,
	}
}

type nameOutcome struct {
	eoutcome.OutcomeBase
	reason int
}

func (z nameOutcome) IsInvalidChar() bool {
	return z.reason == nameOutcomeInvalidChar
}

func (z nameOutcome) IsNameReserved() bool {
	return z.reason == nameOutcomeNameReserved
}

func (z nameOutcome) IsNameTooLong() bool {
	return z.reason == nameOutcomeNameTooLong
}
