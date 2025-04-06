package efs_alpha

import (
	"errors"
)

// Error type constants
var (
	// ErrInvalidChar indicates that the name contains an invalid character
	ErrInvalidChar = errors.New("invalid character in name")

	// ErrNameReserved indicates that the name is reserved
	ErrNameReserved = errors.New("name is reserved")

	// ErrNameTooLong indicates that the name is too long
	ErrNameTooLong = errors.New("name is too long")
)

type Name interface {
	Accept(name string) error
}

// IsInvalidCharError determines if an error indicates an invalid character
// See more detail about limitation:
// https://en.wikipedia.org/wiki/Filename#Comparison_of_filename_limitations
func IsInvalidCharError(err error) bool {
	return err != nil && errors.Is(err, ErrInvalidChar)
}

// IsNameReservedError determines if an error indicates a reserved name
func IsNameReservedError(err error) bool {
	return err != nil && errors.Is(err, ErrNameReserved)
}

// IsNameTooLongError determines if an error indicates a name that is too long
func IsNameTooLongError(err error) bool {
	return err != nil && errors.Is(err, ErrNameTooLong)
}
