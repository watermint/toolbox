package edesktop

import (
	"github.com/watermint/toolbox/essentials/islet/eidiom"
)

type Desktop interface {
	// Open Launches the associated application to open a file or a URL
	Open(p string) OpenOutcome
}

type OpenOutcome interface {
	eidiom.Outcome

	IsOpenFailure() bool
	IsOperationUnsupported() bool
}
