package efs_base

import "fmt"

type Path interface {
	fmt.Stringer

	// Namespace of the path.
	Namespace() Namespace

	// Names returns each path element names from the root.
	Names() []string

	// Basename returns the last element of the path. Returns empty string if the path is "/".
	Basename() string

	// Extname returns the extension of the last element of the path.
	// Returns empty if no ext like ('\.(.)+$') part found.
	Extname() string

	// Root returns root of the file system or namespace.
	Root() Path

	// IsRoot returns true when the path is root path of this file system or namespace.
	IsRoot() bool

	// HasPrefix tests whether the path begins with prefix.
	HasPrefix(prefix string) bool

	// HasSuffix tests whether the path ends with suffix.
	HasSuffix(suffix string) bool

	// Depth returns number of elements
	Depth() int

	// Relative returns a relative path between this path and a given path.
	// Relative does not consider difference in both namespaces.
	// For example, if this path is Windows path `C:\MyData` and
	// the other is UNIX `/MyData/Case01/data001.dat` the relative path is
	// `Case01/data001.dat`
	Relative(other Path) RelativePath

	// Resolve resolves relative path to this path.
	Resolve(relative RelativePath) (Path, PathError)

	// Sibling resolves the given path against this path's parent path.
	// For example, a path represents "dir1/dir2/foo", then invoking this function with the "bar" will result in the Path "dir1/dir2/bar".
	Sibling(other string) (Path, PathError)

	// Parent resolves the parent folder path.
	// Returns root path if the path already root. Never returns nil.
	Parent() Path

	// Child creates child path.
	Child(name ...string) (Path, PathError)
}

type PathError interface {
	error

	// IsPathTooLong returns true when the path exceeds maximum length (entire path or name of entry)
	IsPathTooLong() bool

	// IsPathInvalidName returns true when the path contains invalid character or reserved name
	IsPathInvalidName() bool
}

type RelativePath interface {
	fmt.Stringer

	// Route returns each steps of the path route.
	Route() []PathRoute

	// IsSame returns true when the relative path is identical to the other.
	IsSame() bool

	// IsUpwardOnly returns true when the relative path go towards parent folder(s).
	IsUpwardOnly() bool

	// IsDownwardOnly returns true when the relative path go towards child folder(s)/file(s).
	IsDownwardOnly() bool

	// IsOther returns true when the relative path go towards parent folder(s), then towards child folder(s)/file(s).
	IsOther() bool
}

type PathRoute interface {
	// IsUpward is true when the route is for parent folder.
	IsUpward() bool

	// IsDownward is true when the route is down to child folder.
	IsDownward() bool

	// Name returns path entity name. Returns empty when the route is upward.
	Name() string
}
