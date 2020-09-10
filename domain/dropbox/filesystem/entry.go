package filesystem

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
	"time"
)

func NewEntry(entry mo_file.Entry) es_filesystem.Entry {
	return &dbxEntry{
		entry: entry,
	}
}

type dbxEntry struct {
	entry mo_file.Entry
}

func (z dbxEntry) namespaceId() string {
	ce := z.entry.Concrete()
	if ce.SharedFolderId != "" {
		return ce.SharedFolderId
	} else {
		return ce.ParentSharedFolderId
	}
}

func (z dbxEntry) Name() string {
	return z.entry.Name()
}

func (z dbxEntry) Path() es_filesystem.Path {
	return NewPath(z.namespaceId(), z.entry.Path())
}

func (z dbxEntry) Size() int64 {
	if f, ok := z.entry.File(); ok {
		return f.Size
	}
	return 0
}

func (z dbxEntry) ModTime() time.Time {
	if f, ok := z.entry.File(); ok {
		mt, err := dbx_util.Parse(f.ClientModified)
		if err != nil {
			l := esl.Default()
			l.Warn("Invalid time format", esl.Error(err), esl.Any("entry", z.entry.Concrete()))
			panic(err)
		}
		return mt
	}
	return time.Time{}
}

func (z dbxEntry) ContentHash() (string, es_filesystem.FileSystemError) {
	if f, ok := z.entry.File(); ok {
		return f.ContentHash, nil
	}
	return "", NewError(ErrorInvalidEntryType)
}

func (z dbxEntry) IsFile() bool {
	_, ok := z.entry.File()
	return ok
}

func (z dbxEntry) IsFolder() bool {
	_, ok := z.entry.Folder()
	return ok
}

func (z dbxEntry) AsData() es_filesystem.EntryData {
	return es_filesystem.EntryData{
		FileSystemType: FileSystemTypeDropbox,
		EntryName:      z.Name(),
		EntryPath:      z.Path().Path(),
		EntrySize:      z.Size(),
		EntryModTime:   z.ModTime(),
		EntryIsFile:    z.IsFile(),
		EntryIsFolder:  z.IsFolder(),
		Attributes: map[string]interface{}{
			"raw": z.entry.Concrete().Raw,
		},
	}
}

// Convert ot Dropbox entry model
func ToDropboxEntry(entry es_filesystem.Entry) (mo_file.Entry, error) {
	switch e := entry.(type) {
	case *dbxEntry:
		return e.entry, nil
	default:
		return nil, ErrorInvalidEntryDataFormat
	}
}
