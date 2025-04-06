package es_open

import (
	"errors"
	"fmt"
)

// Error constants
var (
	ErrOpenFailure          = errors.New("failed to open")
	ErrOperationUnsupported = errors.New("operation not supported")
)

// OpenError represents an error that occurred during opening operation
type OpenError struct {
	Err    error
	Reason int
}

const (
	ReasonSuccess = iota
	ReasonOpenFailure
	ReasonOperationUnsupported
)

func (e *OpenError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%v: %v", e.baseError(), e.Err)
	}
	return e.baseError().Error()
}

func (e *OpenError) baseError() error {
	switch e.Reason {
	case ReasonOpenFailure:
		return ErrOpenFailure
	case ReasonOperationUnsupported:
		return ErrOperationUnsupported
	default:
		return nil
	}
}

func (e *OpenError) Is(target error) bool {
	return errors.Is(e.baseError(), target)
}

func (e *OpenError) Unwrap() error {
	return e.Err
}

// IsOpenFailure checks if the error is due to failure to open
func IsOpenFailure(err error) bool {
	var openErr *OpenError
	if errors.As(err, &openErr) {
		return openErr.Reason == ReasonOpenFailure
	}
	return errors.Is(err, ErrOpenFailure)
}

// IsOperationUnsupported checks if the error is due to unsupported operation
func IsOperationUnsupported(err error) bool {
	var openErr *OpenError
	if errors.As(err, &openErr) {
		return openErr.Reason == ReasonOperationUnsupported
	}
	return errors.Is(err, ErrOperationUnsupported)
}

// NewOpenOutcomeSuccess returns nil for successful operation (no error)
func NewOpenOutcomeSuccess() error {
	return nil
}

// NewOpenOutcomeOpenFailure returns an error for open failure
func NewOpenOutcomeOpenFailure(err error) error {
	return &OpenError{
		Err:    err,
		Reason: ReasonOpenFailure,
	}
}

// NewOpenOutcomeOperationUnsupported returns an error for unsupported operation
func NewOpenOutcomeOperationUnsupported(err error) error {
	return &OpenError{
		Err:    err,
		Reason: ReasonOperationUnsupported,
	}
}
