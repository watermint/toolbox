package efs_alpha

import (
	"errors"
)

var (
	// ErrInvalidPath indicates an invalid path
	ErrInvalidPath = errors.New("invalid path")
)

type FileSystem interface {
	Identity() string

	// Path resolves path in the file system. This func does not verify file/folder existence.
	Path(path string) (Path, error)

	// Equals compare to the other. returns true if the other is exactly same as the instance, otherwise false.
	Equals(other FileSystem) bool

	CurrentPath() (Path, error)

	NameRule() Name
}

// IsInvalidPathError determines if an error indicates an invalid path
func IsInvalidPathError(err error) bool {
	return err != nil && errors.Is(err, ErrInvalidPath)
}
