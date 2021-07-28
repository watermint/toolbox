package lang

import "strings"

type MultiError struct {
	allErrors []error
}

func (z MultiError) Error() string {
	errMsgs := make([]string, len(z.allErrors))
	for i, e := range z.allErrors {
		errMsgs[i] = e.Error()
	}
	return "multiple errors: " + strings.Join(errMsgs, ", ")
}

func (z MultiError) NumErrors() int {
	return len(z.allErrors)
}

func (z MultiError) AllErrors() []error {
	return z.allErrors
}

func newMultiError(errors ...error) *MultiError {
	validErrors := make([]error, 0)
	for _, e := range errors {
		if e != nil {
			validErrors = append(validErrors, e)
		}
	}
	return &MultiError{
		allErrors: validErrors,
	}
}

func NewMultiErrorOrNull(errors ...error) error {
	me := newMultiError(errors...)
	switch me.NumErrors() {
	case 0:
		return nil
	case 1:
		return me.AllErrors()[0]
	default:
		return me
	}
}
