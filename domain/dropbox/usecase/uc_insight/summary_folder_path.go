package uc_insight

import "github.com/watermint/toolbox/essentials/log/esl"

type SummaryFolderPath struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`
	FolderId    string `path:"folder_id" gorm:"primaryKey"`

	// parent folder ids joined by slash ('/')
	Path string `path:"path"`

	Updated uint64 `gorm:"autoUpdateTime"`
}

func (z tsImpl) summarizeFolderPaths(folderId string) error {
	l := z.ctl.Log().With(esl.String("folderId", folderId))
	parents := make([]string, 0)
	entry := &NamespaceEntry{}
	if err := z.db.First(entry, "file_id = ?", folderId).Error; err != nil {
		return err
	}
	current := entry.ParentFolderId

	for current != "" {
		f := &NamespaceEntry{}
		if err := z.db.First(f, "file_id = ?", current).Error; err != nil {
			return err
		}
		if f.ParentFolderId != "" {
			parents = append(parents, f.ParentFolderId)
		}
		current = f.ParentFolderId
	}
	path := ""
	for i := len(parents) - 1; i >= 0; i-- {
		path += "/" + parents[i]
	}

	err := z.db.Save(&SummaryFolderPath{
		NamespaceId: entry.NamespaceId,
		FolderId:    folderId,
		Path:        path,
	}).Error
	if err != nil {
		l.Debug("cannot store summary folder path", esl.Error(err), esl.String("namespaceId", entry.NamespaceId), esl.Any("entry", entry))
		return err
	}
	return nil
}
