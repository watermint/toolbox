package mo_file

import "encoding/json"

type Export struct {
	Raw              json.RawMessage
	EntryName        string `path:"file_metadata.name" json:"name"`
	EntryPathLower   string `path:"file_metadata.path_lower" json:"path_lower"`
	EntryPathDisplay string `path:"file_metadata.path_display" json:"path_display"`
	Id               string `path:"file_metadata.id" json:"id"`
	ClientModified   string `path:"file_metadata.client_modified" json:"client_modified"`
	ServerModified   string `path:"file_metadata.server_modified" json:"server_modified"`
	Revision         string `path:"file_metadata.rev" json:"revision"`
	Size             int64  `path:"file_metadata.size" json:"size"`
	ContentHash      string `path:"file_metadata.content_hash" json:"content_hash"`
	ExportName       string `path:"export_metadata.name" json:"export_name"`
	ExportSize       int64  `path:"export_metadata.size" json:"export_size"`
	ExportHash       string `path:"export_metadata.export_hash" json:"export_hash"`
}
