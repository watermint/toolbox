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
	ParentFolderId           string
	EntryType                string `path:"\\.tag"`
	Name                     string `path:"name"`
	Size                     uint64 `path:"size"`
	Rev                      string `path:"rev"`
	IsDownloadable           bool   `path:"is_downloadable"`
	HasExplicitSharedMembers bool   `path:"has_explicit_shared_members"`
	ClientModified           string `path:"client_modified"`
	ServerModified           string `path:"server_modified"`
	ContentHash              string `path:"content_hash"`
	PathLower                string `path:"path_lower"`
	PathDisplay              string `path:"path_display"`
	EntryNamespaceId         string `path:"sharing_info.shared_folder_id"`
	ParentNamespaceId        string `path:"sharing_info.parent_shared_folder_id"`

	Updated uint64 `gorm:"autoUpdateTime"`

	Raw json.RawMessage
}

func NewNamespaceEntry(namespaceId string, parentFolderId string, data es_json.Json) (ne *NamespaceEntry, err error) {
	ne = &NamespaceEntry{}
	if err = data.Model(ne); err != nil {
		return nil, err
	}
	ne.NamespaceId = namespaceId
	ne.ParentFolderId = parentFolderId
	return ne, nil
}
