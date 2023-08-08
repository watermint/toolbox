package uc_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

// SummaryFolderImmediateCount is a summary of the folder.
// This summary is a summary of immediate entries of under the folder, which does not include
// sub folders or those entries.
type SummaryFolderImmediateCount struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`
	FolderId    string `path:"folder_id" gorm:"primaryKey"`

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

func (z tsImpl) summarizeFolderImmediateCount(folder *NamespaceEntry) error {
	rows, err := z.db.Select(&NamespaceEntry{}).Distinct("file_id").Where("parent_folder_id = ?", folder.FileId).Rows()
	if err != nil {
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	fic := &SummaryFolderImmediateCount{
		NamespaceId: folder.NamespaceId,
		FolderId:    folder.FileId,
	}

	for rows.Next() {
		entry := &NamespaceEntry{}
		if err := z.db.ScanRows(rows, entry); err != nil {
			return err
		}
		e := es_json.MustParse(entry.Raw)

		fic.CountEntries++
		if t, ok := e.FindString("\\.tag"); ok {
			switch t {
			case "file":
				fic.CountFile++
				fic.SizeFile += entry.Size
			case "folder":
				fic.CountFolder++
				if entry.EntryNamespaceId != "" {
					fic.CountNamespace++
				}
			case "deleted":
				fic.CountDeleted++
			}
		}
		if t, ok := e.FindBool("is_downloadable"); ok && !t {
			fic.CountNonDownloadable++
		}
	}

	z.db.Save(fic)

	return nil
}
