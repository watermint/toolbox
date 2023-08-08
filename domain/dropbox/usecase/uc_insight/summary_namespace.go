package uc_insight

type SummaryNamespace struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`

	// counts
	CountFile            uint64 `path:"count_file"`
	CountFolder          uint64 `path:"count_folder"`
	CountDeleted         uint64 `path:"count_deleted"`
	CountEntries         uint64 `path:"count_entries"`
	CountNonDownloadable uint64 `path:"count_non_downloadable"`
	CountSymlink         uint64 `path:"count_symlink"`
	CountNamespace       uint64 `path:"count_namespace"`

	// size
	SizeFile uint64 `path:"size_file"`
}
