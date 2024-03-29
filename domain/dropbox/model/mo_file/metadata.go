package mo_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type Metadata struct {
	Raw              json.RawMessage
	EntryTag         string `path:"\\.tag" json:"tag"`
	EntryName        string `path:"name" json:"name"`
	EntryPathDisplay string `path:"path_display" json:"path_display"`
	EntryPathLower   string `path:"path_lower" json:"path_lower"`
}

func (z *Metadata) Metadata() *Metadata {
	return z
}

func (z *Metadata) LockInfo() *LockInfo {
	return newLockInfo(z.Raw)
}

func (z *Metadata) Tag() string {
	return z.EntryTag
}

func (z *Metadata) Name() string {
	return z.EntryName
}

func (z *Metadata) PathDisplay() string {
	return z.EntryPathDisplay
}

func (z *Metadata) PathLower() string {
	return z.EntryPathLower
}

func (z *Metadata) Path() mo_path.DropboxPath {
	return mo_path.NewPathDisplay(z.EntryPathDisplay)
}

func (z *Metadata) File() (*File, bool) {
	if z.EntryTag != "file" {
		return nil, false
	}
	f := &File{}
	if err := api_parser.ParseModelRaw(f, z.Raw); err != nil {
		return nil, false // Should not happen
	}
	return f, true
}

func (z *Metadata) Folder() (*Folder, bool) {
	if z.EntryTag != "folder" {
		return nil, false
	}
	f := &Folder{}
	if err := api_parser.ParseModelRaw(f, z.Raw); err != nil {
		return nil, false
	}
	return f, true
}

func (z *Metadata) Deleted() (*Deleted, bool) {
	if z.EntryTag != "deleted" {
		return nil, false
	}
	d := &Deleted{}
	if err := api_parser.ParseModelRaw(d, z.Raw); err != nil {
		return nil, false
	}
	return d, true
}

func (z *Metadata) Concrete() *ConcreteEntry {
	return newConcreteEntry(z.Raw)
}

func newMetadataEntry(raw json.RawMessage) (entry *Metadata) {
	entry = &Metadata{}
	if err := api_parser.ParseModelRaw(entry, raw); err != nil {
		esl.Default().Debug("Unable to parse json", esl.Error(err), esl.ByteString("raw", raw))
		return entry
	}
	entry.Raw = raw
	return entry
}
