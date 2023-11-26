package es_open

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated/eoutcome"
)

const (
	openOutcomeSuccess = iota
	openOutcomeOpenFailure
	openOutcomeOperationUnsupported
)

func NewOpenOutcomeSuccess() OpenOutcome {
	return &openOutcomeImpl{
		OutcomeBase: eoutcome.NewOutcomeBaseOk(),
		reason:      openOutcomeSuccess,
	}
}

func NewOpenOutcomeOpenFailure(err error) OpenOutcome {
	return &openOutcomeImpl{
		OutcomeBase: eoutcome.NewOutcomeBaseError(err),
		reason:      openOutcomeOpenFailure,
	}
}

func NewOpenOutcomeOperationUnsupported(err error) OpenOutcome {
	return &openOutcomeImpl{
		OutcomeBase: eoutcome.NewOutcomeBaseError(err),
		reason:      openOutcomeOperationUnsupported,
	}
}

type openOutcomeImpl struct {
	eoutcome.OutcomeBase
	reason int
}

func (z openOutcomeImpl) IsOpenFailure() bool {
	return z.reason == openOutcomeOpenFailure
}

func (z openOutcomeImpl) IsOperationUnsupported() bool {
	return z.reason == openOutcomeOperationUnsupported
}
