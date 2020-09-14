package es_filesystem_local

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"os"
)

func NewEntry(path string, info os.FileInfo) es_filesystem.Entry {
	return &fsEntry{
		EntryData: es_filesystem.EntryData{
			FileSystemType: FileSystemTypeLocal,
			EntryName:      info.Name(),
			EntryPath:      path,
			EntrySize:      info.Size(),
			EntryModTime:   info.ModTime(),
			EntryIsFile:    !info.IsDir(),
			EntryIsFolder:  info.IsDir(),
		},
	}
}

type fsEntry struct {
	es_filesystem.EntryData
}

func (z fsEntry) Path() es_filesystem.Path {
	return NewPath(z.EntryPath)
}

func (z fsEntry) ContentHash() (string, es_filesystem.FileSystemError) {
	hash, err := dbx_util.FileContentHash(z.EntryPath)
	return hash, NewError(err)
}
