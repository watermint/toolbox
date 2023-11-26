package es_open

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated"
)

type Desktop interface {
	// Open Launches the associated application to open a file or a URL
	Open(p string) OpenOutcome
}

type OpenOutcome interface {
	es_idiom_deprecated.Outcome

	IsOpenFailure() bool
	IsOperationUnsupported() bool
}
