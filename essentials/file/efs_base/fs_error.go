package efs_base

type ErrorBase interface {
	error

	// Cause returns original error
	Cause() error
}

type FsError interface {
	ErrorBase
	PathError

	// IsNotAllowed returns true in case the caller does not allowed to execute the operation.
	// POSIX file systems: true when a caller does not allowed to access local file system.
	// Cloud file systems: true when a caller have valid OAuth access token but the token does not issued with required scopes.
	IsNotAllowed() bool

	// IsTimeout returns true when the operation timeout.
	IsTimeout() bool

	// IsConflict returns true if an operation failed due to conflict with existing file or folder.
	// But the operation should not return an error if the result is the same as what was requested.
	// For example, if a caller called CreateFolder and the folder already exist in the path specified.
	// That is identical to requested state. The function should not return an error this case.
	IsConflict() bool

	// IsPermission returns true if an operation failed due to missing permission to the path.
	// For example, when a caller called PutFile to the folder and the caller does not have permission to write that folder.
	IsPermission() bool
}
