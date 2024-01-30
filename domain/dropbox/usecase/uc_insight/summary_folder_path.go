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
	Name        string

	// parent folder ids joined by slash ('/')
	Path string `path:"path"`
	// PathDisplay is Display path from top of the namespace
	PathDisplay string `path:"path_display"`

	Updated uint64 `gorm:"autoUpdateTime"`
}

func (z tsImpl) summarizeFolderPaths(folderId string) error {
	l := z.ctl.Log().With(esl.String("folderId", folderId))
	entry := &NamespaceEntry{}
	entryPath := ""
	entryNames := make([]string, 0)

	parents := make([]string, 0)
	if err := z.adb.First(entry, "file_id = ?", folderId).Error; err != nil {
		l.Debug("cannot retrieve entry", esl.Error(err), esl.String("folderId", folderId))
		return err
	}
	current := entry.ParentFolderId

	ns := &Namespace{}
	if entry.EntryNamespaceId != "" {
		if err := z.adb.First(ns, "namespace_id = ?", entry.EntryNamespaceId).Error; err != nil {
			l.Debug("cannot retrieve namespace", esl.Error(err), esl.String("namespaceId", entry.EntryNamespaceId))
			// fall through
		}
	}

	for current != "" {
		ne := &NamespaceEntry{}
		if err := z.adb.First(ne, "file_id = ?", current).Error; err != nil {
			// This is not unusual. The cause hypothesis: in case the parent folder is not visible to the user
			l.Debug("cannot retrieve parent entry", esl.Error(err), esl.String("current", current))
			break
		}
		if ne.ParentFolderId != "" {
			parents = append(parents, ne.ParentFolderId)
		}

		cn := &Namespace{}
		if err := z.adb.First(cn, "namespace_id = ?", ne.NamespaceId).Error; err != nil {
			l.Debug("cannot retrieve namespace", esl.Error(err), esl.String("namespaceId", ne.NamespaceId))
			return err
		}

		if cn.NamespaceType != "team_member_folder" {
			entryNames = append(entryNames, ne.Name)
		}
		current = ne.ParentFolderId
	}
	slices.Reverse(entryNames)
	if ns.NamespaceType != "team_member_folder" {
		entryNames = append(entryNames, entry.Name)
	}
	path := ""
	for i := len(parents) - 1; i >= 0; i-- {
		path += "/" + parents[i]
	}
	entryPath = path

	err := z.sdb.Save(&SummaryFolderPath{
		FolderId:    folderId,
		Name:        entry.Name,
		NamespaceId: entry.NamespaceId,
		Path:        entryPath,
		PathDisplay: strings.Join(entryNames, "/"),
	}).Error
	if err != nil {
		l.Debug("cannot store summary folder path", esl.Error(err), esl.String("namespaceId", entry.NamespaceId), esl.Any("entry", entry))
		return err
	}

	return nil
}
