package efs_deprecated

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated"
)

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
	es_idiom_deprecated.Outcome

	IsInvalidPath() bool
}

type CurrentPathOutcome interface {
	es_idiom_deprecated.Outcome
}
