package uc_insight

type SummaryCount struct {
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

func (z SummaryCount) AddEntry(entry *NamespaceEntry) SummaryCount {
	z.CountEntries++
	switch entry.EntryType {
	case "file":
		z.CountFile++
		z.SizeFile += entry.Size
		if !entry.IsDownloadable {
			z.CountNonDownloadable++
		}
	case "folder":
		z.CountFolder++
		if entry.EntryNamespaceId != "" {
			z.CountNamespace++
		}
	case "deleted":
		z.CountDeleted++
	}
	return z
}

func (z SummaryCount) AddSummary(y SummaryCount) SummaryCount {
	z.CountFile += y.CountFile
	z.CountFolder += y.CountFolder
	z.CountDeleted += y.CountDeleted
	z.CountEntries += y.CountEntries
	z.CountNonDownloadable += y.CountNonDownloadable
	z.CountSymlink += y.CountSymlink
	z.CountNamespace += y.CountNamespace
	z.SizeFile += y.SizeFile
	return z
}
