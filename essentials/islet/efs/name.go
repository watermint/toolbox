package efs

import "github.com/watermint/toolbox/essentials/islet/eidiom"

type Name interface {
	Accept(name string) NameOutcome
}

type NameOutcome interface {
	eidiom.Outcome

	// IsInvalidChar return true if invalid char found in given name
	// See more detail about limitation:
	// https://en.wikipedia.org/wiki/Filename#Comparison_of_filename_limitations
	IsInvalidChar() bool

	// IsNameReserved IsReserved return true if the name is reserved by the system
	IsNameReserved() bool
	IsNameTooLong() bool
}
