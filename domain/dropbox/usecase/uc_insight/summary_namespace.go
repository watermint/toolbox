package uc_insight

import (
	"github.com/watermint/toolbox/essentials/log/esl"
)

type SummaryNamespace struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`

	RootNamespaceType string
	RootNamespaceId   string

	// counts
	SummaryCount

	Updated uint64 `gorm:"autoUpdateTime"`
}

func (z tsImpl) summarizeNamespace(namespaceId string) error {
	l := z.ctl.Log().With(esl.String("namespaceId", namespaceId))
	sn := &SummaryNamespace{
		NamespaceId: namespaceId,
	}
	sc := SummaryCount{}

	rootNamespaceId := ""
	currentNamespaceId := namespaceId
	for {
		nd := &NamespaceDetail{}
		if err := z.adb.First(nd, "namespace_id = ?", currentNamespaceId).Error; err != nil {
			ns := &Namespace{}
			if err := z.adb.First(ns, "namespace_id = ?", currentNamespaceId).Error; err != nil {
				l.Debug("cannot find namespace", esl.Error(err))
				return err
			}
			rootNamespaceId = ns.NamespaceId
			break
		}
		currentNamespaceId = nd.ParentNamespaceId
		if currentNamespaceId == "" {
			rootNamespaceId = nd.NamespaceId
			break
		}
	}

	rootNamespace := &Namespace{}
	if err := z.adb.First(rootNamespace, "namespace_id = ?", rootNamespaceId).Error; err != nil {
		l.Debug("cannot find namespace", esl.Error(err))
		return err
	}

	sn.RootNamespaceId = rootNamespaceId
	sn.RootNamespaceType = rootNamespace.NamespaceType

	rows, err := z.adb.Model(&NamespaceEntry{}).Where("namespace_id = ?", namespaceId).Rows()
	if err != nil {
		l.Debug("cannot find entries", esl.Error(err))
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		entry := &NamespaceEntry{}
		if err := z.adb.ScanRows(rows, entry); err != nil {
			l.Debug("cannot scan row", esl.Error(err))
			return err
		}
		sc = sc.AddEntry(entry)
	}

	sn.SummaryCount = sc
	if err := z.sdb.Save(&sn).Error; err != nil {
		l.Debug("cannot save", esl.Error(err))
		return err
	}
	return nil

}
