package efs_deprecated

import "fmt"

const (
	EntryTypeFolder = iota
	EntryTypeFile
	EntryTypeSymlink
	EntryTypeDeleted
	EntryTypeDeletedFile
	EntryTypeDeletedFolder
	EntryTypeOnlineDocument
	EntryTypeReparsePointsSymlink
	EntryTypePlaceholder
)

// Entry is an element for representation of child item of a folder, that is typically a file or a folder.
type Entry interface {
	fmt.Stringer

	// Name base name of the entry.
	Name() string
}
