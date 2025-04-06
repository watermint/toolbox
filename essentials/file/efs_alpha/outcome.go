package efs_alpha

import (
	"errors"
)

var (
	// ErrTimeout indicates a timeout error
	ErrTimeout = errors.New("operation timed out")

	// ErrOperationNotAllowed indicates an operation is not allowed
	ErrOperationNotAllowed = errors.New("operation not allowed")
)

// IsTimeoutError determines if an error indicates a timeout
func IsTimeoutError(err error) bool {
	return err != nil && errors.Is(err, ErrTimeout)
}

// IsOperationNotAllowedError determines if an error indicates an operation is not allowed
func IsOperationNotAllowedError(err error) bool {
	return err != nil && errors.Is(err, ErrOperationNotAllowed)
}
