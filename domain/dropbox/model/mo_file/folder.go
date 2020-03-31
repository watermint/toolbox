package mo_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
)

type Folder struct {
	Raw                       json.RawMessage
	Id                        string `path:"id" json:"id"`
	EntryTag                  string `path:"\\.tag" json:"tag"`
	EntryName                 string `path:"name" json:"name"`
	EntryPathLower            string `path:"path_lower" json:"path_lower"`
	EntryPathDisplay          string `path:"path_display" json:"path_display"`
	EntrySharedFolderId       string `path:"sharing_info.shared_folder_id" json:"shared_folder_id"`
	EntryParentSharedFolderId string `path:"sharing_info.parent_shared_folder_id" json:"parent_shared_folder_id"`
}

func (z *Folder) Tag() string {
	return z.EntryTag
}

func (z *Folder) Name() string {
	return z.EntryName
}

func (z *Folder) PathDisplay() string {
	return z.EntryPathDisplay
}

func (z *Folder) PathLower() string {
	return z.EntryPathLower
}

func (z *Folder) Path() mo_path.DropboxPath {
	return mo_path.NewPathDisplay(z.EntryPathDisplay)
}

func (z *Folder) File() (*File, bool) {
	return nil, false
}

func (z *Folder) Folder() (*Folder, bool) {
	return z, true
}

func (z *Folder) Deleted() (*Deleted, bool) {
	return nil, false
}

func (z *Folder) SharedFolderId() string {
	return z.EntrySharedFolderId
}

func (z *Folder) ParentSharedFolderId() string {
	return z.EntryParentSharedFolderId
}

func (z *Folder) Concrete() *ConcreteEntry {
	return newConcreteEntry(z.Raw)
}
