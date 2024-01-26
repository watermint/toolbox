package uc_insight

import (
	"database/sql"
	"github.com/watermint/toolbox/essentials/log/esl"
	"gorm.io/gorm"
)

// SummaryFolderImmediateCount is a summary of the folder.
// This summary is a summary of immediate entries of under the folder, which does not include
// sub folders or those entries.
type SummaryFolderImmediateCount struct {
	// primary keys
	FolderId string `path:"folder_id" gorm:"primaryKey"`

	SummaryCount

	Updated uint64 `gorm:"autoUpdateTime"`
}

type SummaryFolderAndNamespace struct {
	// primary keys
	FolderId string `path:"folder_id" gorm:"primaryKey"`

	SummaryCount
}

func (z tsImpl) summarizeFolderImmediateCount(folderId string) error {
	l := z.ctl.Log().With(esl.String("folderId", folderId))
	if folderId == "" {
		l.Debug("skip. no folder id")
		return nil
	}
	summaryImmediate := SummaryCount{}
	summaryFolderAndNamespace := SummaryCount{}

	err := z.db.Transaction(func(tx *gorm.DB) error {
		ne := &NamespaceEntry{}
		rows, err := tx.Model(ne).Where("parent_folder_id = ?", folderId).Rows()
		if err != nil {
			l.Debug("cannot find entries", esl.Error(err))
			return err
		}
		defer func() {
			_ = rows.Close()
		}()

		for rows.Next() {
			entry := &NamespaceEntry{}
			if err := tx.ScanRows(rows, entry); err != nil {
				return err
			}
			summaryImmediate = summaryImmediate.AddEntry(entry)

			if entry.EntryType == "folder" && entry.EntryNamespaceId != "" {
				ns := &SummaryNamespace{}
				if err := tx.First(ns, "namespace_id = ?", entry.EntryNamespaceId).Error; err != nil {
					l.Debug("cannot find namespace", esl.Error(err), esl.String("namespaceId", entry.EntryNamespaceId))
					return err
				}
				summaryFolderAndNamespace = summaryFolderAndNamespace.AddSummary(ns.SummaryCount)
			}
		}

		return nil
	}, &sql.TxOptions{
		ReadOnly: true,
	})
	if err != nil {
		l.Debug("cannot summarize", esl.Error(err))
		return err
	}

	z.db.Save(&SummaryFolderImmediateCount{
		FolderId:     folderId,
		SummaryCount: summaryImmediate,
	})
	z.db.Save(&SummaryFolderAndNamespace{
		FolderId:     folderId,
		SummaryCount: summaryFolderAndNamespace.AddSummary(summaryImmediate),
	})

	return nil
}
