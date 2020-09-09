package es_filesystem_model

import "github.com/watermint/toolbox/essentials/file/es_filesystem"

const (
	ErrorTypePathNotFound = iota
	ErrorTypeConflict
	ErrorTypeNoPermission
	ErrorTypeInsufficientSpace
	ErrorTypeDisallowedName
	ErrorTypeInvalidEntryDataFormat
	ErrorTypeOther
)

func NewError(err error, errType int) es_filesystem.FileSystemError {
	if err == nil {
		return nil
	}
	return &FileSystemError{
		err: err,
	}
}

type FileSystemError struct {
	err     error
	errType int
}

func (z FileSystemError) Error() string {
	return z.err.Error()
}

func (z FileSystemError) IsPathNotFound() bool {
	return z.errType == ErrorTypePathNotFound
}

func (z FileSystemError) IsConflict() bool {
	return z.errType == ErrorTypeConflict
}

func (z FileSystemError) IsNoPermission() bool {
	return z.errType == ErrorTypeNoPermission
}

func (z FileSystemError) IsInsufficientSpace() bool {
	return z.errType == ErrorTypeInsufficientSpace
}

func (z FileSystemError) IsDisallowedName() bool {
	return z.errType == ErrorTypeDisallowedName
}

func (z FileSystemError) IsInvalidEntryDataFormat() bool {
	return z.errType == ErrorTypeInvalidEntryDataFormat
}
