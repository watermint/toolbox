package efs_base

import (
	"io"
	"time"
)

type PathBaseOps interface {
	// Path parses string
	Path(path string, names ...string) (Path, PathError)
}

type FileBaseOps interface {
	FilePut(path Path, data io.Reader) FsError
	FileGet(path Path) (data io.ReadCloser, fsErr FsError)
	FileSize(path Path) (size int64, fsError FsError)
	FileCopy(src, dst Path) FsError
	FileMove(src, dst Path) FsError
	FileDelete(path Path) FsError
}

type FolderEntry interface {
	Path() Path
	IsFile() bool
	IsFolder() bool
}

type FolderBaseOps interface {
	FolderCreate(path Path) FsError
	FolderList(path Path, onEntry func(entry FolderEntry) bool) FsError
	FolderCopy(src, dst Path) FsError
	FolderMove(src, dst Path) FsError
	FolderDelete(path Path) FsError
}

type StateBaseOps interface {
	IsExist(path Path) (bool, FsError)
}

type TypeBaseOps interface {
	IsFolder(path Path) (bool, FsError)
	IsFile(path Path) (bool, FsError)
}

type TimeBaseOps interface {
	TimeModifiedGet(path Path) (time.Time, FsError)
}

type FileSystemBase interface {
	PathBaseOps
	FileBaseOps
	FolderBaseOps
	TypeBaseOps
	StateBaseOps
	TimeBaseOps
}
