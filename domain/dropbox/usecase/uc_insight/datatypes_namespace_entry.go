package uc_insight

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type NamespaceEntry struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`
	FileId      string `path:"id" gorm:"primaryKey"`

	// attributes
	EntryType                string `path:".tag"`
	Name                     string `path:"name"`
	Size                     uint64 `path:"size"`
	Rev                      string `path:"rev"`
	IsDownloadable           bool   `path:"is_downloadable"`
	HasExplicitSharedMembers bool   `path:"has_explicit_shared_members"`
	ClientModified           string `path:"client_modified"`
	ServerModified           string `path:"server_modified"`
	Path                     string `path:"path"`
	PathLower                string `path:"path_lower"`
	PathDisplay              string `path:"path_display"`
	ParentSharedFolderId     string `path:"sharing_info.parent_shared_folder_id"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

func NewNamespaceEntry(namespaceId string, path string, data es_json.Json) (ne *NamespaceEntry, err error) {
	ne = &NamespaceEntry{}
	if err = data.Model(ne); err != nil {
		return nil, err
	}
	ne.NamespaceId = namespaceId
	ne.Path = path
	return ne, nil
}
