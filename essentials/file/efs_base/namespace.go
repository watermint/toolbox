package efs_base

import "fmt"

// Namespace is the local disk volume, remote hostname mount point, or cloud storage namespace.
//
// Windows example; Returns "C:" for "C:\foo\bar", "\\host\share" for "\\host\share\foo.txt", or
// "\\?\C:" for "\\?\C:\foo.txt".
// UN*X example; Returns "" for "/foo/bar", or "host:" for "host:/foo/bar".
// Cloud storage example; returns namespace id like "ns:123456789".
type Namespace interface {
	fmt.Stringer

	// IsEmpty returns true for default local file system or no namespace defined in the file system.
	IsEmpty() bool
}
