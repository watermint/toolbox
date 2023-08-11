package uc_insight

import "github.com/watermint/toolbox/essentials/log/esl"

type SummaryNamespace struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`

	// counts
	SummaryCount

	Updated uint64 `gorm:"autoUpdateTime"`
}

func (z tsImpl) summarizeNamespace(namespaceId string) error {
	l := z.ctl.Log().With(esl.String("namespaceId", namespaceId))
	rows, err := z.db.Model(&NamespaceEntry{}).Where("namespace_id = ?", namespaceId).Rows()
	if err != nil {
		l.Debug("cannot find entries", esl.Error(err))
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	sc := SummaryCount{}
	for rows.Next() {
		entry := &NamespaceEntry{}
		if err := z.db.ScanRows(rows, entry); err != nil {
			l.Debug("cannot scan row", esl.Error(err))
			return err
		}
		sc = sc.AddEntry(entry)
	}
	z.db.Save(&SummaryNamespace{
		NamespaceId:  namespaceId,
		SummaryCount: sc,
	})
	return nil
}
