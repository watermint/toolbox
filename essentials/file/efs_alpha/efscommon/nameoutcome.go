package efscommon

import (
	"fmt"

	"github.com/watermint/toolbox/essentials/file/efs_alpha"
)

const (
	nameOutcomeSuccess = iota
	nameOutcomeInvalidChar
	nameOutcomeNameReserved
	nameOutcomeNameTooLong
)

// NewNameOutcomeSuccess returns nil to represent success
func NewNameOutcomeSuccess() error {
	return nil
}

// NewNameOutcomeInvalidChar returns an invalid character error
func NewNameOutcomeInvalidChar(invalidChar string) error {
	return fmt.Errorf("%w: invalid char '%s' found", efs_alpha.ErrInvalidChar, invalidChar)
}

// NewNameOutcomeNameReserved returns a reserved name error
func NewNameOutcomeNameReserved(reserved string) error {
	return fmt.Errorf("%w: reserved keyword '%s' found", efs_alpha.ErrNameReserved, reserved)
}

// NewNameOutcomeNameTooLong returns a name too long error
func NewNameOutcomeNameTooLong(length, max int) error {
	return fmt.Errorf("%w: name is too long (%d). maximum length is %d", efs_alpha.ErrNameTooLong, length, max)
}
