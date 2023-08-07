package uc_insight

type NamespaceEntryError struct {
	NamespaceId string `path:"shared_folder_id" gorm:"primaryKey"`
	FolderId    string `path:"folder_id" gorm:"primaryKey"`

	Error string `path:"error_summary"`

	Updated uint64 `gorm:"autoUpdateTime"`
}
