package eoutcome

import (
	"errors"
	"github.com/watermint/toolbox/essentials/islet/eidiom"
)

type ParseOutcome interface {
	eidiom.Outcome

	IsInvalidFormat() bool
	IsInvalidChar() bool
}

const (
	parseReasonSuccess = iota
	parseReasonInvalidFormat
	parseReasonInvalidChar
)

func NewParseSuccess() ParseOutcome {
	return &parseOutcomeImpl{
		OutcomeBase: OutcomeBase{Err: nil},
		reasonCode:  parseReasonSuccess,
	}
}

func NewParseInvalidFormat(reason string) ParseOutcome {
	return &parseOutcomeImpl{
		OutcomeBase:  OutcomeBase{Err: errors.New(reason)},
		reasonCode:   parseReasonInvalidFormat,
		reasonDetail: reason,
	}
}

func NewParseInvalidChar(reason string) ParseOutcome {
	return &parseOutcomeImpl{
		OutcomeBase:  OutcomeBase{Err: errors.New(reason)},
		reasonCode:   parseReasonInvalidChar,
		reasonDetail: reason,
	}
}

type parseOutcomeImpl struct {
	OutcomeBase
	reasonCode   int
	reasonDetail string
}

func (z parseOutcomeImpl) IsInvalidChar() bool {
	return z.reasonCode == parseReasonInvalidChar
}

func (z parseOutcomeImpl) IsInvalidFormat() bool {
	return z.reasonCode == parseReasonInvalidFormat
}
