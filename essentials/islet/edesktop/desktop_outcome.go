package edesktop

import "github.com/watermint/toolbox/essentials/islet/eidiom/eoutcome"

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
