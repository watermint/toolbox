package efs

import "github.com/watermint/toolbox/essentials/islet/eidiom"

type FileSystem interface {
	Identity() string

	// Path resolves path in the file system. This func does not verify file/folder existence.
	Path(path string) (Path, PathOutcome)

	// Equals compare to the other. returns true if the other is exactly same as the instance, otherwise false.
	Equals(other FileSystem) bool

	CurrentPath() (Path, CurrentPathOutcome)

	NameRule() Name
}

type PathOutcome interface {
	eidiom.Outcome

	IsInvalidPath() bool
}

type CurrentPathOutcome interface {
	eidiom.Outcome
}
