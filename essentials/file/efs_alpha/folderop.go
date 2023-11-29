package efs_alpha

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated"
)

type FolderOps interface {
	ListEntries(f func(entry Entry) es_idiom_deprecated.Outcome) ListOutcome
	ListFiles(f func(file File) es_idiom_deprecated.Outcome) ListOutcome
	ListFolders(f func(folder Folder) es_idiom_deprecated.Outcome) ListOutcome

	WalkEntries(f func(entry Entry) es_idiom_deprecated.Outcome) WalkOutcome
	WalkFiles(f func(file File) es_idiom_deprecated.Outcome) WalkOutcome
	WalkFolders(f func(folder Folder) es_idiom_deprecated.Outcome) WalkOutcome
}

type ListOutcome interface {
	FileSystemOutcome
}

type WalkOutcome interface {
	FileSystemOutcome
}
