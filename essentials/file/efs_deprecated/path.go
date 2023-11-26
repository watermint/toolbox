package efs_deprecated

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated"
)

// Path abstract absolute file path for various file systems.
type Path interface {
	fmt.Stringer

	// Identity is for identify path. It designed for compare paths with identity whether it equals to another.
	// The identity string does not need to be portable.
	// It might point to the another path if some context change, e.g. different user context.
	// The path implementation must use path separator for '/' even for the Windows implementation.
	//
	// Case sensitivity:
	// The identity string should be lowercase for case in-sensitive file systems.
	// Changing case should use user's context such as locale,
	// for example, in case of changing case from "I" to "i", change to "ı" (dotless I) for Turkish locale.
	// https://en.wikipedia.org/wiki/Dotted_and_dotless_I
	Identity() string

	// FileSystem returns associated file system instance.
	FileSystem() FileSystem

	// Parent returns parent path.
	Parent() Path

	// Basename Returns the last element of the path. Returns empty string if the path is "/".
	Basename() string

	// Extname returns the extension of the last element of the path.
	// Returns empty if no ext like ('\.(.)+$') part found.
	Extname() string

	// IsRoot returns true when the path is root path of this file system or namespace.
	IsRoot() bool

	// Namespace of the path. Returns empty string if no namespace associated to the path.
	//
	// Windows example; Returns "C:" for "C:\foo\bar", "\\host\share" for "\\host\share\foo.txt", or
	// "\\?\C:" for "\\?\C:\foo.txt".
	// UN*X example; Returns "" for "/foo/bar", or "host:" for "host:/foo/bar".
	Namespace() Namespace

	// Child creates child path.
	Child(name ...string) (Path, ChildOutcome)
}

type AbsoluteOutcome interface {
	es_idiom_deprecated.Outcome

	IsFileSystemNotAccessible() bool
	IsNoDefaultPath() bool
}

type ChildOutcome interface {
	NameOutcome

	// IsPathTooLong returns true if the path is too long for the system.
	IsPathTooLong() bool
}
