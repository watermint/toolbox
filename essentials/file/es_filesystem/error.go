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
}
