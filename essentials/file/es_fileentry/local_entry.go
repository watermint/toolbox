package es_fileentry

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

// os.FileInfo equivalent serializable structure
type LocalEntry struct {
	// fully qualified path
	Path string `json:"path"`

	// name of the entry
	Name string `json:"name"`

	// File size
	Size int64 `json:"size"`

	// Modification time
	ModTime time.Time `json:"mod_time"`

	// File mode
	Mode os.FileMode `json:"mode"`

	// Is Folder
	IsFolder bool
}

func NewLocalEntry(path string, entry os.FileInfo) LocalEntry {
	return LocalEntry{
		Path:     filepath.Join(path, entry.Name()),
		Name:     entry.Name(),
		Size:     entry.Size(),
		ModTime:  entry.ModTime(),
		Mode:     entry.Mode(),
		IsFolder: entry.IsDir(),
	}
}

func ReadLocalEntries(path string) (entries []LocalEntry, err error) {
	osEntries, err := os.ReadDir(path)
	if err != nil {
		return
	}
	entries = make([]LocalEntry, len(osEntries))
	for i, osEntry := range osEntries {
		info, err := osEntry.Info()
		if err != nil {
			return nil, err
		}
		entries[i] = NewLocalEntry(path, info)
	}
	return
}

func LocalEntryByNameLower(entries []LocalEntry) (byName map[string]LocalEntry) {
	byName = make(map[string]LocalEntry)
	for _, entry := range entries {
		byName[strings.ToLower(entry.Name)] = entry
	}
	return
}
