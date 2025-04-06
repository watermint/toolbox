package efs_alpha

import (
	"errors"
)

// FileOps defines operation set for a file.
type FileOps interface {
	CreateFile() (File, error)
	DeleteFile() error
	DeleteFileIfExists() error
}

var (
	// ErrFileNotFound indicates a file was not found
	ErrFileNotFound = errors.New("file not found")
)

// IsFileNotFoundError determines if an error indicates a file was not found
func IsFileNotFoundError(err error) bool {
	return err != nil && errors.Is(err, ErrFileNotFound)
}
