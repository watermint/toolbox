package efs_deprecated

type PathOps interface {
	Exist() (bool, FileSystemOutcome)
	Move(dest Path) MoveOutcome
	Copy(dest Path) CopyOutcome
}

type MoveOutcome interface {
	FileSystemOutcome
}

type CopyOutcome interface {
	FileSystemOutcome
}
