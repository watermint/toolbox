package efs

import "github.com/watermint/toolbox/essentials/islet/eidiom"

type FolderOps interface {
	ListEntries(f func(entry Entry) eidiom.Outcome) ListOutcome
	ListFiles(f func(file File) eidiom.Outcome) ListOutcome
	ListFolders(f func(folder Folder) eidiom.Outcome) ListOutcome

	WalkEntries(f func(entry Entry) eidiom.Outcome) WalkOutcome
	WalkFiles(f func(file File) eidiom.Outcome) WalkOutcome
	WalkFolders(f func(folder Folder) eidiom.Outcome) WalkOutcome
}

type ListOutcome interface {
	FileSystemOutcome
}

type WalkOutcome interface {
	FileSystemOutcome
}
