package uc_insight

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"golang.org/x/exp/slices"
	"strings"
)

type SummaryFolderPath struct {
	// primary keys
	NamespaceId string `path:"namespace_id" gorm:"primaryKey"`
	FolderId    string `path:"folder_id" gorm:"primaryKey"`

	// parent folder ids joined by slash ('/')
	Path string `path:"path"`
	// PathDisplay is Display path from top of the namespace
	PathDisplay string `path:"path_display"`

	Updated uint64 `gorm:"autoUpdateTime"`
}

func (z tsImpl) summarizeFolderPaths(folderId string) error {
	l := z.ctl.Log().With(esl.String("folderId", folderId))
	parents := make([]string, 0)
	entry := &NamespaceEntry{}
	if err := z.db.First(entry, "file_id = ?", folderId).Error; err != nil {
		l.Debug("cannot retrieve entry", esl.Error(err), esl.String("folderId", folderId))
		return err
	}
	current := entry.ParentFolderId

	ns := &Namespace{}
	if err := z.db.First(ns, "namespace_id = ?", entry.NamespaceId).Error; err != nil {
		l.Debug("cannot retrieve namespace", esl.Error(err), esl.String("namespaceId", entry.NamespaceId))
		return err
	}

	names := make([]string, 0)
	for current != "" {
		ne := &NamespaceEntry{}
		if err := z.db.First(ne, "file_id = ?", current).Error; err != nil {
			l.Debug("cannot retrieve entry", esl.Error(err), esl.String("folderId", current))
			return err
		}
		if ne.ParentFolderId != "" {
			parents = append(parents, ne.ParentFolderId)
		}

		cn := &Namespace{}
		if err := z.db.First(cn, "namespace_id = ?", ne.NamespaceId).Error; err != nil {
			l.Debug("cannot retrieve namespace", esl.Error(err), esl.String("namespaceId", ne.NamespaceId))
			return err
		}

		if cn.NamespaceType != "team_member_folder" {
			names = append(names, ne.Name)
		}
		current = ne.ParentFolderId
	}
	slices.Reverse(names)
	if ns.NamespaceType != "team_member_folder" {
		names = append(names, ns.Name)
	}
	path := ""
	for i := len(parents) - 1; i >= 0; i-- {
		path += "/" + parents[i]
	}

	err := z.db.Save(&SummaryFolderPath{
		NamespaceId: entry.NamespaceId,
		FolderId:    folderId,
		Path:        path,
		PathDisplay: strings.Join(names, "/"),
	}).Error
	if err != nil {
		l.Debug("cannot store summary folder path", esl.Error(err), esl.String("namespaceId", entry.NamespaceId), esl.Any("entry", entry))
		return err
	}
	return nil
}
