package es_filesystem_local

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	FileSystemTypeLocal = "local"
)

func NewFileSystem() es_filesystem.FileSystem {
	return fsLocal{}
}

type fsLocal struct {
}

func (z fsLocal) OperationalComplexity(entries []es_filesystem.Entry) (complexity int64) {
	return int64(len(entries))
}

func (z fsLocal) Path(data es_filesystem.PathData) (path es_filesystem.Path, err es_filesystem.FileSystemError) {
	if data.FileSystemType != FileSystemTypeLocal {
		return nil, NewError(ErrorInvalidEntryDataFormat)
	}

	return &fsPath{path: data.EntryPath}, nil
}

func (z fsLocal) Shard(data es_filesystem.ShardData) (namespace es_filesystem.Shard, err es_filesystem.FileSystemError) {
	if data.FileSystemType != FileSystemTypeLocal {
		return nil, NewError(ErrorInvalidEntryDataFormat)
	}

	return data, nil
}

func (z fsLocal) Entry(data es_filesystem.EntryData) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	if data.FileSystemType != FileSystemTypeLocal {
		return nil, NewError(ErrorInvalidEntryDataFormat)
	}

	return &fsEntry{EntryData: data}, nil
}

func (z fsLocal) FileSystemType() string {
	return FileSystemTypeLocal
}

func (z fsLocal) List(path es_filesystem.Path) (entries []es_filesystem.Entry, err es_filesystem.FileSystemError) {
	osEntries, osErr := ioutil.ReadDir(path.Path())
	if osErr != nil {
		return nil, NewError(osErr)
	}
	entries = make([]es_filesystem.Entry, 0)
	for _, osEntry := range osEntries {
		entries = append(entries, NewEntry(filepath.Join(path.Path(), osEntry.Name()), osEntry))
	}
	return
}

func (z fsLocal) Info(path es_filesystem.Path) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	osEntry, osErr := os.Lstat(path.Path())
	if osErr != nil {
		return nil, NewError(osErr)
	}
	return NewEntry(path.Path(), osEntry), nil
}

func (z fsLocal) Delete(path es_filesystem.Path) (err es_filesystem.FileSystemError) {
	osErr := os.RemoveAll(path.Path())
	return NewError(osErr)
}

func (z fsLocal) CreateFolder(path es_filesystem.Path) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	osErr := os.MkdirAll(path.Path(), 0755)
	if osErr != nil {
		return nil, NewError(osErr)
	}
	return z.Info(path)
}
