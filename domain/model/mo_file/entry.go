package mo_file

import "github.com/watermint/toolbox/domain/model/mo_path"

type Entry interface {
	// Tag for the entry. `file`, `folder`, or `deleted`.
	Tag() string

	// File or folder basename
	Name() string

	// Display path
	PathDisplay() string

	// Lowercase path
	PathLower() string

	// Path
	Path() mo_path.Path

	// Returns File, returns nil & false if an entry is not a File.
	File() (*File, bool)

	// Returns Folder, return nil & false if an entry is not a Folder.
	Folder() (*Folder, bool)

	// Returns Deleted, return nil & false if an entry is not a Deleted.
	Deleted() (*Deleted, bool)
}
