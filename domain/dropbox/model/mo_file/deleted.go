package mo_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
)

type Deleted struct {
	Raw              json.RawMessage
	EntryTag         string `path:"\\.tag" json:"tag"`
	EntryName        string `path:"name" json:"name"`
	EntryPathLower   string `path:"path_lower" json:"path_lower"`
	EntryPathDisplay string `path:"path_display" json:"path_display"`
}

func (z *Deleted) LockInfo() *LockInfo {
	return nil
}

func (z *Deleted) Tag() string {
	return z.EntryTag
}

func (z *Deleted) Name() string {
	return z.EntryName
}

func (z *Deleted) PathDisplay() string {
	return z.EntryPathDisplay
}

func (z *Deleted) PathLower() string {
	return z.EntryPathLower
}

func (z *Deleted) Path() mo_path.DropboxPath {
	return mo_path.NewPathDisplay(z.EntryPathDisplay)
}

func (z *Deleted) File() (*File, bool) {
	return nil, false
}

func (z *Deleted) Folder() (*Folder, bool) {
	return nil, false
}

func (z *Deleted) Deleted() (*Deleted, bool) {
	return z, true
}

func (z *Deleted) Concrete() *ConcreteEntry {
	return newConcreteEntry(z.Raw)
}
