package es_errors

import (
	"errors"
	"fmt"
)

// Standard error constants
const (
	// ParseErrorPrefix is the prefix for parse-related errors
	ParseErrorPrefix = "parse error: "

	// InvalidFormat indicates invalid format
	InvalidFormat = "invalid format"

	// OutOfRange indicates a value outside the allowed range
	OutOfRange = "out of range"

	// Reserved indicates a reserved item
	Reserved = "reserved"

	// TooLong indicates that something is too long
	TooLong = "too long"

	// InvalidChar indicates an invalid character
	InvalidChar = "invalid character"
)

// Standard error type definitions
var (
	// ErrParse is a general parsing error
	ErrParse = errors.New(ParseErrorPrefix + "failed")

	// ErrInvalidFormat indicates an invalid format
	ErrInvalidFormat = errors.New(ParseErrorPrefix + InvalidFormat)

	// ErrOutOfRange indicates a value that is out of range
	ErrOutOfRange = errors.New(ParseErrorPrefix + OutOfRange)
)

// NewParseError creates a parse-related error
func NewParseError(format string, args ...interface{}) error {
	return fmt.Errorf(ParseErrorPrefix+format, args...)
}

// IsParseError determines if an error is a parse-related error
func IsParseError(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, ErrParse) ||
		errors.Is(err, ErrInvalidFormat) ||
		errors.Is(err, ErrOutOfRange) ||
		(len(err.Error()) >= len(ParseErrorPrefix) &&
			err.Error()[:len(ParseErrorPrefix)] == ParseErrorPrefix)
}

// NewInvalidFormatError creates an invalid format error
func NewInvalidFormatError(format string, args ...interface{}) error {
	if len(args) == 0 {
		return ErrInvalidFormat
	}
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s: %s", ErrInvalidFormat.Error(), msg)
}

// IsInvalidFormatError determines if an error is an invalid format error
func IsInvalidFormatError(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, ErrInvalidFormat) ||
		(len(err.Error()) >= len(ParseErrorPrefix+InvalidFormat) &&
			err.Error()[:len(ParseErrorPrefix+InvalidFormat)] == ParseErrorPrefix+InvalidFormat)
}

// NewOutOfRangeError creates an out-of-range error
func NewOutOfRangeError(format string, args ...interface{}) error {
	if len(args) == 0 {
		return ErrOutOfRange
	}
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s: %s", ErrOutOfRange.Error(), msg)
}

// IsOutOfRangeError determines if an error is an out-of-range error
func IsOutOfRangeError(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, ErrOutOfRange) ||
		(len(err.Error()) >= len(ParseErrorPrefix+OutOfRange) &&
			err.Error()[:len(ParseErrorPrefix+OutOfRange)] == ParseErrorPrefix+OutOfRange)
}
