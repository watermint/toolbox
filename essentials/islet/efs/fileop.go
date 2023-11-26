package efs

// FileOps defines operation set for a file.
type FileOps interface {
	CreateFile() (File, CreateFileOutcome)
	DeleteFile() DeleteFileOutcome
	DeleteFileIfExists() DeleteFileOutcome
}

type CreateFileOutcome interface {
	FileSystemOutcome
}

type DeleteFileOutcome interface {
	FileSystemOutcome

	IsFileNotFound() bool
}
