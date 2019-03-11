package mo_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/model/mo_path"
)

type File struct {
	Raw              json.RawMessage
	EntryTag         string `path:"\\.tag"`
	EntryName        string `path:"name"`
	EntryPathLower   string `path:"path_lower"`
	EntryPathDisplay string `path:"path_display"`
	Id               string `path:"id"`
	ClientModified   string `path:"client_modified"`
	ServerModified   string `path:"server_modified"`
	Revision         string `path:"rev"`
	Size             int64  `path:"size"`
	ContentHash      string `path:"content_hash"`
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

func (z *File) Path() mo_path.Path {
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
