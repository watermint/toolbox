package efs_alpha

type FolderOps interface {
	ListEntries(f func(entry Entry) error) error
	ListFiles(f func(file File) error) error
	ListFolders(f func(folder Folder) error) error

	WalkEntries(f func(entry Entry) error) error
	WalkFiles(f func(file File) error) error
	WalkFolders(f func(folder Folder) error) error
}
