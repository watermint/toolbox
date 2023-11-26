package eoutcome

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/essentials/islet/eidiom"
)

func NewOutcomeBaseOk() OutcomeBase {
	return OutcomeBase{
		Err: nil,
	}
}

func NewOutcomeBaseError(err error) OutcomeBase {
	return OutcomeBase{
		Err: err,
	}
}

func NewOutcomeBaseWithErrMessage(format string, v ...interface{}) OutcomeBase {
	return OutcomeBase{
		Err: errors.New(fmt.Sprintf(format, v...)),
	}
}

type OutcomeBase struct {
	Err error
}

func (z OutcomeBase) String() string {
	if z.Err != nil {
		errStr := z.Err.Error()
		if errStr == "" {
			return "error"
		}
		return errStr
	} else {
		return ""
	}
}

func (z OutcomeBase) Cause() error {
	return z.Err
}

func (z OutcomeBase) IsOk() bool {
	return z.Err == nil
}

func (z OutcomeBase) IfOk(f func()) {
	if z.Err == nil {
		f()
	}
}

func (z OutcomeBase) IsError() bool {
	return z.Err != nil
}

func (z OutcomeBase) IfError(f func() eidiom.Outcome) eidiom.Outcome {
	if z.Err != nil {
		return f()
	}
	return z
}

// AssertOutcomeSuccess test the outcome implementation conforms behavior for succeed
func AssertOutcomeSuccess(outcome eidiom.Outcome) bool {
	if outcome.IsError() {
		return false
	}
	if !outcome.IsOk() {
		return false
	}
	if outcome.String() != "" {
		return false
	}
	if outcome.Cause() != nil {
		return false
	}
	ifErr := false
	outcome.IfError(func() eidiom.Outcome {
		ifErr = true
		return outcome
	})
	if ifErr {
		return false
	}
	ifOk := false
	outcome.IfOk(func() {
		ifOk = true
	})
	if !ifOk {
		return false
	}
	return true
}

// AssertOutcomeFailure test the outcome implementation conforms behavior for error
func AssertOutcomeFailure(outcome eidiom.Outcome) bool {
	if !outcome.IsError() {
		return false
	}
	if outcome.IsOk() {
		return false
	}
	if outcome.String() == "" {
		return false
	}
	if outcome.Cause() == nil {
		return false
	}
	ifErr := false
	outcome.IfError(func() eidiom.Outcome {
		ifErr = true
		return outcome
	})
	if !ifErr {
		return false
	}
	ifOk := false
	outcome.IfOk(func() {
		ifOk = true
	})
	if ifOk {
		return false
	}
	return true
}
