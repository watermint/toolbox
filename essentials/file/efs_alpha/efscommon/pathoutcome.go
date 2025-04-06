package efscommon

import (
	"fmt"

	"github.com/watermint/toolbox/essentials/file/efs_alpha"
)

// NewChildOutcomeSuccess returns nil to represent success
func NewChildOutcomeSuccess() error {
	return nil
}

// NewChildOutcomeByNameOutcome returns an error based on name validation
func NewChildOutcomeByNameOutcome(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("invalid name: %w", err)
}

// NewChildOutcomePathTooLong returns an error indicating that the path is too long
func NewChildOutcomePathTooLong(given, allowed int) error {
	return fmt.Errorf("%w: given length (%d), allowed length (%d)", efs_alpha.ErrPathTooLong, given, allowed)
}
