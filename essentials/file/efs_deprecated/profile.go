package efs_deprecated

// Profile collection of information about file system with current principle.
// All profiles are knowledge based on the file system implementation.
// Some attribute may not true for a specific path. For example,
// On UNIX, base file system usually supports read/write.
// But you can mount readonly device such as CD/DVD to a certain mount point.
type Profile interface {
	// IsPosixPermission returns true if the file system uses POSIX compatible permission model, otherwise false.
	IsPosixPermission() bool

	// IsHardLinkSupported returns true if the file system supports hard link, otherwise false.
	IsHardLinkSupported() bool

	// IsSymbolicLinkSupported returns true if the file system supports symbolic link, otherwise false.
	IsSymbolicLinkSupported() bool

	// IsReparsePointsSupported returns true if the file system supports reparse points, otherwise false.
	IsReparsePointsSupported() bool

	// IsOnlineDocumentSupported returns true if the file systems supports special online document such as Google Docs, otherwise false.
	IsOnlineDocumentSupported() bool

	// IsFileLockSupported returns true if the file system supports lock/unlock a file, otherwise false.
	IsFileLockSupported() bool

	// IsFileCommentSupported returns true if the file system supports comment for a file, otherwise false.
	IsFileCommentSupported() bool

	// IsFileVersionHistorySupported returns true if the file system supports file version history, otherwise false.
	IsFileVersionHistorySupported() bool

	// IsHiddenAttrSupported returns true if the file system supports hidden file, otherwise false.
	IsHiddenAttrSupported() bool

	// IsReadOnly returns true if the file system is read only, otherwise false.
	IsReadOnly() bool

	// MaxFileSize returns maximum file size.
	MaxFileSize() (uint64, ProfileOutcome)

	// MaxNumFiles returns maximum number of files per single namespace.
	MaxNumFiles(namespace Namespace) (uint64, ProfileOutcome)
}

type ProfileOutcome interface {
	FileSystemOutcome

	IsUnknown() bool
}
