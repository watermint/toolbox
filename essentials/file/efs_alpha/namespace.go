package efs_alpha

import "fmt"

type Namespace interface {
	fmt.Stringer

	// Identity The path implementation must use path separator for '/' even for the Windows implementation.
	// The identity string does not need to be portable.
	// It might point to the another path if some context change, e.g. different user context.
	//
	// Case sensitivity: see more detail at Path#Identity
	Identity() string

	// Equals compare to the other. returns true if the other is exactly same as this instance, otherwise false.
	Equals(other Namespace) bool

	// IsDefault returns true if the namespace is the default namespace of current file system and principle.
	IsDefault() bool

	// FileSystem returns the file system.
	FileSystem() FileSystem
}
