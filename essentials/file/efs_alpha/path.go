package efs_alpha

import (
	"errors"
	"fmt"
)

// Path is an interface to handle paths
type Path interface {
	fmt.Stringer

	// IsRoot returns true if this path is a root path.
	IsRoot() bool

	// Name returns the name of the path. Returns empty string when path is root.
	Name() string

	// Extension returns the extension, including dot (e.g.) ".txt"
	Extension() string

	// Basename returns the base name of the path.
	Basename() string

	// DirSlash returns the directory portion with trailing slash.
	DirSlash() string

	// Dir returns the directory portion.
	Dir() Path

	// Child creates a new path from the given relative path.
	Child(path string) (Path, error)

	// FindFile recursive searches a file that match with the condition.
	// Return nil if not match.
	FindFile(cond func(path Path) bool) (Path, error)

	// FindPathUnderTree retrieves all paths under path tree (search recursively) that matches of the condition.
	// The result array sorted by path name. Empty array if no match.
	FindPathUnderTree(cond func(path Path) bool) ([]Path, error)

	// Walk traverses filesystem hierarchy under the path.
	// The result array sorted by path name.
	Walk(fn func(path Path) error) error

	// MustChild is same as Child but returns empty path if error.
	MustChild(path string) Path

	// RelativeTo returns a relative path that is lexically equivalent to targpath when joined to basepath.
	// Returns an error if targpath can't be made relative to basepath.
	RelativeTo(base Path) (string, error)

	// Equal checks file system path equality
	Equal(p Path) bool
}

// ChildPathType PATH vs CHILD types
type ChildPathType int

const (
	// ChildPathTypeUnknown
	ChildPathTypeUnknown ChildPathType = iota

	// ChildPathTypeParentRelative relative path like `../folder`
	ChildPathTypeParentRelative

	// ChildPathTypeAbsolute absolute path like `/foo/bar`, `C:\Users`
	ChildPathTypeAbsolute

	// ChildPathTypeRelative relative path like `foo/bar/baz`
	ChildPathTypeRelative
)

// AbsoluteOutcome interface
type AbsoluteOutcome interface {
	Cause() error
}

var (
	// ErrPathTooLong indicates that the path is too long
	ErrPathTooLong = errors.New("path too long")
)

// IsPathTooLongError determines if an error indicates a path that is too long
func IsPathTooLongError(err error) bool {
	return err != nil && errors.Is(err, ErrPathTooLong)
}
