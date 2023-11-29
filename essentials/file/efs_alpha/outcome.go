package efs_alpha

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated"
)

type FileSystemOutcome interface {
	es_idiom_deprecated.Outcome

	// IsTimeout returns true if an operation failed with timeout.
	IsTimeout() bool

	// IsOperationNotAllowed returns true if an operation is not allowed.
	IsOperationNotAllowed() bool
}
