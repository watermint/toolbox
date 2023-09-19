package mo_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"github.com/watermint/toolbox/essentials/log/esl"
	"strings"
)

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
	Path() mo_path.DropboxPath

	// Returns File, returns nil & false if an entry is not a File.
	File() (*File, bool)

	// Returns Folder, return nil & false if an entry is not a Folder.
	Folder() (*Folder, bool)

	// Returns Deleted, return nil & false if an entry is not a Deleted.
	Deleted() (*Deleted, bool)

	// Returns metadata instance.
	Metadata() *Metadata

	// Returns concrete entry
	Concrete() *ConcreteEntry

	// Returns lock info, returns nil if the entry does not have lock info (e.g. a deleted entry or a folder)
	LockInfo() *LockInfo
}

type ConcreteEntry struct {
	Raw                      json.RawMessage
	Id                       string `path:"id" json:"id"`
	Tag                      string `path:"\\.tag" json:"tag"`
	Name                     string `path:"name" json:"name"`
	PathLower                string `path:"path_lower" json:"path_lower"`
	PathDisplay              string `path:"path_display" json:"path_display"`
	ClientModified           string `path:"client_modified" json:"client_modified"`
	ServerModified           string `path:"server_modified" json:"server_modified"`
	Revision                 string `path:"rev" json:"revision"`
	Size                     int64  `path:"size" json:"size"`
	ContentHash              string `path:"content_hash" json:"content_hash"`
	HasExplicitSharedMembers bool   `path:"has_explicit_shared_members" json:"has_explicit_shared_members"`
	SharedFolderId           string `path:"sharing_info.shared_folder_id" json:"shared_folder_id"`
	ParentSharedFolderId     string `path:"sharing_info.parent_shared_folder_id" json:"parent_shared_folder_id"`
}

func (z ConcreteEntry) IsFile() bool {
	return z.Tag == "file"
}

func (z ConcreteEntry) IsFolder() bool {
	return z.Tag == "folder"
}

func (z ConcreteEntry) IsDeleted() bool {
	return z.Tag == "deleted"
}

func newConcreteEntry(raw json.RawMessage) *ConcreteEntry {
	ce := &ConcreteEntry{}
	if err := api_parser.ParseModelRaw(ce, raw); err != nil {
		esl.Default().Debug("Unable to parse json", esl.Error(err), esl.ByteString("raw", raw))
		return ce
	}
	ce.Raw = raw
	return ce
}

func MapByNameLower(entries []Entry) map[string]Entry {
	mte := make(map[string]Entry)
	for _, entry := range entries {
		mte[strings.ToLower(entry.Name())] = entry
	}
	return mte
}
