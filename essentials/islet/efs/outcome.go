package efs

import "github.com/watermint/toolbox/essentials/islet/eidiom"

type FileSystemOutcome interface {
	eidiom.Outcome

	// IsTimeout returns true if an operation failed with timeout.
	IsTimeout() bool

	// IsOperationNotAllowed returns true if an operation is not allowed.
	IsOperationNotAllowed() bool
}
