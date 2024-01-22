package uc_insight

type SummaryTeamFolderEntry struct {
	// primary keys
	TeamFolderId string `path:"team_folder_id" gorm:"primaryKey"`
	FileId       string `path:"file_id" gorm:"primaryKey"`

	Name             string `path:"name"`
	EntryType        string `path:"entry_type"`
	ParentFolderId   string `path:"parent_folder_id" gorm:"index"`
	EntryNamespaceId string `path:"entry_namespace_id" gorm:"index"`

	PathDisplay string `path:"path_display"`

	// Updated is the timestamp when the entry is updated.
	Updated uint64 `gorm:"autoUpdateTime"`
}
