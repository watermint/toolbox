package uc_insight

import "github.com/watermint/toolbox/essentials/log/esl"

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
	ne := &NamespaceEntry{}
	rows, err := z.db.Model(ne).Where("parent_folder_id = ?", folderId).Rows()
	if err != nil {
		l.Debug("cannot find entries", esl.Error(err))
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	sc := SummaryCount{}
	sn := SummaryCount{}

	for rows.Next() {
		entry := &NamespaceEntry{}
		if err := z.db.ScanRows(rows, entry); err != nil {
			return err
		}
		sc = sc.AddEntry(entry)

		if entry.EntryType == "folder" && entry.EntryNamespaceId != "" {
			ns := &SummaryNamespace{}
			if err := z.db.First(ns, "namespace_id = ?", entry.EntryNamespaceId).Error; err != nil {
				l.Debug("cannot find namespace", esl.Error(err), esl.String("namespaceId", entry.EntryNamespaceId))
				return err
			}
			sn = sn.AddSummary(ns.SummaryCount)
		}
	}

	sfn := sn.AddSummary(sc)
	z.db.Save(&SummaryFolderImmediateCount{
		FolderId:     folderId,
		SummaryCount: sc,
	})
	z.db.Save(&SummaryFolderAndNamespace{
		FolderId:     folderId,
		SummaryCount: sfn,
	})

	return nil
}
