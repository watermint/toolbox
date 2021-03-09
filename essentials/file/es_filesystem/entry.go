package es_filesystem

import "time"

type Entry interface {
	// Name of the entry.
	Name() string

	// Path of the entry.
	Path() Path

	// Size of the entry.
	Size() int64

	// Modification time.
	ModTime() time.Time

	// Content hash.
	ContentHash() (string, FileSystemError)

	// True when the entry is a file.
	IsFile() bool

	// True when the entry is a folder.
	IsFolder() bool

	// Convert as serializable data
	AsData() EntryData
}

func CompareEntry(a, b Entry) bool {
	if a.Name() != b.Name() {
		return false
	}
	if a.IsFile() != b.IsFile() {
		return false
	}
	if a.IsFolder() != b.IsFolder() {
		return false
	}
	if a.Path().Path() != b.Path().Path() {
		return false
	}

	// compare as file
	if a.IsFile() {
		if a.Size() != b.Size() {
			return false
		}
		if a.ModTime() != b.ModTime() {
			return false
		}
		ah, _ := a.ContentHash()
		bh, _ := b.ContentHash()
		if ah != bh {
			return false
		}
	}

	return true
}

type EntryData struct {
	FileSystemType string                 `json:"file_system_type"`
	EntryName      string                 `json:"name"`
	EntryPath      string                 `json:"path"`
	EntrySize      int64                  `json:"size"`
	EntryModTime   time.Time              `json:"mod_time"`
	EntryIsFile    bool                   `json:"is_file"`
	EntryIsFolder  bool                   `json:"is_folder"`
	Attributes     map[string]interface{} `json:"attributes,omitempty"`
}

func (z EntryData) Name() string {
	return z.EntryName
}

func (z EntryData) Size() int64 {
	return z.EntrySize
}

func (z EntryData) ModTime() time.Time {
	return z.EntryModTime
}

func (z EntryData) IsFile() bool {
	return z.EntryIsFile
}

func (z EntryData) IsFolder() bool {
	return z.EntryIsFolder
}

func (z EntryData) AsData() EntryData {
	return z
}
