package es_filesystem

import (
	"path/filepath"
	"time"
)

const (
	FileSystemTypeMock = "mock"
)

func NewMockPath(path string) Path {
	return &MockPath{
		PathData: PathData{
			FileSystemType: FileSystemTypeMock,
			EntryPath:      filepath.ToSlash(filepath.Clean(path)),
			EntryShard:     ShardData{},
			Attributes:     map[string]interface{}{},
		},
	}
}

type MockPath struct {
	PathData
}

func (z MockPath) Base() string {
	return filepath.Base(z.EntryPath)
}

func (z MockPath) Ancestor() Path {
	if z.IsRoot() {
		return z
	} else {
		return NewMockPath(filepath.Dir(z.Path()))
	}
}

func (z MockPath) Descendant(pathFragment ...string) Path {
	fragments := make([]string, 0)
	fragments = append(fragments, z.Path())
	fragments = append(fragments, pathFragment...)
	return NewMockPath(filepath.Join(fragments...))
}

func (z MockPath) IsRoot() bool {
	switch z.Path() {
	case "/", "":
		return true
	default:
		return false
	}
}

func NewMockFileEntry(path string, size int64, modTime time.Time, mockHash string) Entry {
	return &MockEntry{
		EntryData: EntryData{
			FileSystemType: FileSystemTypeMock,
			EntryName:      filepath.Base(path),
			EntryPath:      filepath.ToSlash(filepath.Clean(path)),
			EntrySize:      size,
			EntryModTime:   modTime,
			EntryIsFile:    true,
			EntryIsFolder:  false,
			Attributes:     map[string]interface{}{},
		},
		MockHash: mockHash,
	}
}

func NewMockFolderEntry(path string) Entry {
	return &MockEntry{
		EntryData: EntryData{
			FileSystemType: FileSystemTypeMock,
			EntryName:      filepath.Base(path),
			EntryPath:      filepath.ToSlash(filepath.Clean(path)),
			EntrySize:      0,
			EntryModTime:   time.Time{},
			EntryIsFile:    false,
			EntryIsFolder:  true,
			Attributes:     map[string]interface{}{},
		},
		MockHash: "",
	}
}

type MockEntry struct {
	EntryData
	MockHash string `json:"content_hash"`
}

func (z MockEntry) Path() Path {
	return NewMockPath(z.EntryPath)
}

func (z MockEntry) ContentHash() (string, FileSystemError) {
	return z.MockHash, nil
}
