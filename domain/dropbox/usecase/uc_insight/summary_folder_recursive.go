package uc_insight

import (
	"github.com/watermint/toolbox/essentials/log/esl"
)

type SummaryFolderRecursive struct {
	// primary keys
	FolderId string `path:"folder_id" gorm:"primaryKey"`

	SummaryCount

	Updated uint64 `gorm:"autoUpdateTime"`
}

func (z tsImpl) summarizeFolderRecursive(folderId string) error {
	l := z.ctl.Log().With(esl.String("folderId", folderId))
	sc := SummaryCount{}

	children, err := z.sdb.Model(&SummaryFolderPath{}).Where("path LIKE ?", "%"+folderId+"%").Rows()
	if err != nil {
		l.Debug("cannot count", esl.Error(err))
		return err
	}
	defer func() {
		_ = children.Close()
	}()

	for children.Next() {
		child := &SummaryFolderPath{}
		if err := z.sdb.ScanRows(children, child); err != nil {
			l.Debug("cannot scan child row", esl.Error(err))
			return err
		}
		sn := &SummaryFolderAndNamespace{}
		if err := z.sdb.First(sn, "folder_id = ?", child.FolderId).Error; err != nil {
			l.Debug("cannot find folder", esl.Error(err), esl.String("folderId", child.FolderId))
			return err
		}
		sc = sc.AddSummary(sn.SummaryCount)
		sc.CountFolder++
		sc.CountEntries++
	}

	z.sdb.Save(&SummaryFolderRecursive{
		FolderId:     folderId,
		SummaryCount: sc,
	})
	return nil
}
