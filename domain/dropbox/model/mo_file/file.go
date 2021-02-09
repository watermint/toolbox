package mo_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
)

type File struct {
	Raw              json.RawMessage
	EntryTag         string `path:"\\.tag" json:"tag"`
	EntryName        string `path:"name" json:"name"`
	EntryPathLower   string `path:"path_lower" json:"path_lower"`
	EntryPathDisplay string `path:"path_display" json:"path_display"`
	Id               string `path:"id" json:"id"`
	ClientModified   string `path:"client_modified" json:"client_modified"`
	ServerModified   string `path:"server_modified" json:"server_modified"`
	Revision         string `path:"rev" json:"revision"`
	Size             int64  `path:"size" json:"size"`
	ContentHash      string `path:"content_hash" json:"content_hash"`
}

func (z *File) LockInfo() *LockInfo {
	return newLockInfo(z.Raw)
}

func (z *File) Tag() string {
	return z.EntryTag
}

func (z *File) Name() string {
	return z.EntryName
}

func (z *File) PathDisplay() string {
	return z.EntryPathDisplay
}

func (z *File) PathLower() string {
	return z.EntryPathLower
}

func (z *File) Path() mo_path.DropboxPath {
	return mo_path.NewPathDisplay(z.EntryPathDisplay)
}

func (z *File) File() (*File, bool) {
	return z, true
}

func (z *File) Folder() (*Folder, bool) {
	return nil, false
}

func (z *File) Deleted() (*Deleted, bool) {
	return nil, false
}

func (z *File) Concrete() *ConcreteEntry {
	return newConcreteEntry(z.Raw)
}
