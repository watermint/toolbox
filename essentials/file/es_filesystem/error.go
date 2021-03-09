package es_filesystem

type FileSystemError interface {
	error

	// True if a path is not found.
	IsPathNotFound() bool

	// True if an operation detect conflict of the path.
	IsConflict() bool

	// True if the user doesn't have permissions to write to the target.
	IsNoPermission() bool

	// True if the file system doesn't have enough available space to write more data.
	IsInsufficientSpace() bool

	// True if the file system doesn't allow name of the entry.
	IsDisallowedName() bool

	// True if the entry data is not valid for this file system.
	IsInvalidEntryDataFormat() bool

	// True if the error is for testing
	IsMockError() bool
}

func NewLowLevelError(ge error) FileSystemError {
	return &generalError{
		err: ge,
	}
}

type generalError struct {
	err error
}

func (z generalError) Error() string {
	return z.err.Error()
}

func (z generalError) IsPathNotFound() bool {
	return false
}

func (z generalError) IsConflict() bool {
	return false
}

func (z generalError) IsNoPermission() bool {
	return false
}

func (z generalError) IsInsufficientSpace() bool {
	return false
}

func (z generalError) IsDisallowedName() bool {
	return false
}

func (z generalError) IsInvalidEntryDataFormat() bool {
	return false
}

func (z generalError) IsMockError() bool {
	return false
}
