package eg_image

import (
	"errors"
	"fmt"
)

var (
	// ErrUnsupportedFormatError is an error for unsupported formats
	ErrUnsupportedFormatError = errors.New(ErrUnsupportedFormat)

	// ErrEncodeFailureError is an encoding error
	ErrEncodeFailureError = errors.New(ErrEncodeFailure)

	// ErrWriteFailureError is a writing error
	ErrWriteFailureError = errors.New(ErrWriteFailure)
)

// NewExportOutcomeSuccess returns nil to represent success
func NewExportOutcomeSuccess() error {
	return nil
}

// NewExportOutcomeUnsupportedFormat returns an unsupported format error
func NewExportOutcomeUnsupportedFormat(given ImageFormat) error {
	return fmt.Errorf("%w: unsupported format %d", ErrUnsupportedFormatError, given)
}

// NewExportOutcomeEncodeFailure returns an encoding error
func NewExportOutcomeEncodeFailure(err error) error {
	return fmt.Errorf("%w: %v", ErrEncodeFailureError, err)
}

// NewExportOutcomeWriteFailure returns a writing error
func NewExportOutcomeWriteFailure(err error) error {
	return fmt.Errorf("%w: %v", ErrWriteFailureError, err)
}

// IsUnsupportedFormatError determines if an error is an unsupported format error
func IsUnsupportedFormatError(err error) bool {
	return err != nil && errors.Is(err, ErrUnsupportedFormatError)
}

// IsEncodeFailureError determines if an error is an encoding error
func IsEncodeFailureError(err error) bool {
	return err != nil && errors.Is(err, ErrEncodeFailureError)
}

// IsWriteFailureError determines if an error is a writing error
func IsWriteFailureError(err error) bool {
	return err != nil && errors.Is(err, ErrWriteFailureError)
}
