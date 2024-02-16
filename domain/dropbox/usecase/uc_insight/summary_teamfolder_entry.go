package uc_insight

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"gorm.io/gorm"
	"strings"
)

type SummaryTeamFolderEntry struct {
	// primary keys
	TeamFolderId string `path:"team_folder_id" gorm:"primaryKey"`
	FileId       string `path:"file_id" gorm:"primaryKey"`

	Name             string `path:"name"`
	EntryType        string `path:"entry_type"`
	ParentFolderId   string `path:"parent_folder_id" gorm:"index"`
	EntryNamespaceId string `path:"entry_namespace_id" gorm:"index"`

	// Path is the path concatenated by fileId with slash ('/').
	Path string `path:"path"`
	// PathDisplay is Display path from top of the team folder.
	PathDisplay string `path:"path_display"`

	SummaryMemberCount
	SummaryEntryCount

	// Size
	Size uint64 `path:"size"`

	// Updated is the timestamp when the entry is updated.
	Updated uint64 `gorm:"autoUpdateTime"`
}

type SummarizeTeamFolderEntry struct {
	TeamFolder  *TeamFolder `path:"team_folder" json:"team_folder"`
	FileId      string      `path:"file_id" json:"file_id"`
	ParentId    string      `path:"parent_id" json:"parent_id"`
	ParentIds   []string    `path:"parent_ids" json:"parent_ids"`
	ParentNames []string    `path:"parent_names" json:"parent_names"`
}

func (z summaryImpl) summarizeTeamFolder(teamFolder *TeamFolder, s eq_sequence.Stage) error {
	l := z.ctl.Log().With(esl.String("teamFolderId", teamFolder.TeamFolderId))
	childEntries := make([]*SummarizeTeamFolderEntry, 0)

	ne := &NamespaceEntry{}
	if err := z.db.First(ne, "entry_namespace_id = ?", teamFolder.TeamFolderId).Error; err != nil {
		l.Debug("Unable to find namespace entry", esl.Error(err))
		return err
	}
	if ne.FileId == "" {
		l.Debug("No file id", esl.Any("ne", ne))
		return nil
	}

	rows, err := z.db.Model(&NamespaceEntry{}).Where("parent_folder_id = ? AND entry_type = 'folder'", ne.FileId).Rows()
	if err != nil {
		l.Debug("Unable to find namespace entry", esl.Error(err))
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		cne := &NamespaceEntry{}
		if err := z.db.ScanRows(rows, cne); err != nil {
			l.Debug("Unable to find namespace entry", esl.Error(err))
			return err
		}
		childEntries = append(childEntries, &SummarizeTeamFolderEntry{
			TeamFolder:  teamFolder,
			FileId:      cne.FileId,
			ParentId:    ne.FileId,
			ParentIds:   []string{ne.FileId},
			ParentNames: []string{ne.Name},
		})
	}

	if len(childEntries) > 0 {
		q := s.Get(teamSummarizeTeamFolderEntry)
		for _, ce := range childEntries {
			q.Enqueue(ce)
		}
	}

	return nil
}

func (z summaryImpl) summarizeTeamFolderEntry(teamFolderEntry *SummarizeTeamFolderEntry, s eq_sequence.Stage) error {
	l := z.ctl.Log().With(esl.String("teamFolderId", teamFolderEntry.TeamFolder.TeamFolderId))
	childEntries := make([]*SummarizeTeamFolderEntry, 0)
	if teamFolderEntry.FileId == "" {
		l.Debug("No file id", esl.Any("teamFolderEntry", teamFolderEntry))
		return nil
	}

	ne := &NamespaceEntry{}
	if err := z.db.First(ne, "file_id = ?", teamFolderEntry.FileId).Error; err != nil {
		l.Debug("Unable to find namespace entry", esl.Error(err))
		return err
	}
	se := &SummaryEntry{}
	if ne.EntryType == "folder" {
		err := z.db.First(se, "file_id = ?", teamFolderEntry.FileId).Error
		switch {
		case err == nil, errors.Is(err, gorm.ErrRecordNotFound):
			l.Debug("Processing folder entry")

		default:
			l.Debug("Unable to find summary entry", esl.Error(err))
			return err
		}
	}

	tfe := &SummaryTeamFolderEntry{
		TeamFolderId:     teamFolderEntry.TeamFolder.TeamFolderId,
		FileId:           teamFolderEntry.FileId,
		Name:             ne.Name,
		EntryType:        ne.EntryType,
		ParentFolderId:   teamFolderEntry.ParentId,
		EntryNamespaceId: ne.EntryNamespaceId,
		Path:             strings.Join(append(teamFolderEntry.ParentIds, teamFolderEntry.FileId), "/"),
		PathDisplay:      strings.Join(append(teamFolderEntry.ParentNames, ne.Name), "/"),
		SummaryMemberCount: SummaryMemberCount{
			CountUniqueLower:    se.CountUniqueLower,
			CountUniqueUpper:    se.CountUniqueUpper,
			CountMember:         se.CountMember,
			CountMemberInternal: se.CountMemberInternal,
			CountMemberExternal: se.CountMemberExternal,
			CountInvitee:        se.CountInvitee,
			CountGroup:          se.CountGroup,
			CountGroupExternal:  se.CountGroupExternal,
		},
		Size: se.Size,
	}
	z.db.Save(tfe)

	if ne.EntryType == "folder" && ne.FileId != "" {
		rows, err := z.db.Model(&NamespaceEntry{}).Where("parent_folder_id = ? AND entry_type = 'folder'", ne.FileId).Rows()
		switch {
		case err == nil:
			defer func() {
				_ = rows.Close()
			}()

			for rows.Next() {
				cne := &NamespaceEntry{}
				if err := z.db.ScanRows(rows, cne); err != nil {
					l.Debug("Unable to find namespace entry", esl.Error(err))
					return err
				}

				if cne.FileId == "" {
					l.Debug("No file id", esl.Any("cne", cne))
					z.db.Save(&SummaryTeamFolderEntry{
						TeamFolderId:       teamFolderEntry.TeamFolder.TeamFolderId,
						FileId:             "",
						Name:               cne.Name,
						EntryType:          cne.EntryType,
						ParentFolderId:     cne.ParentFolderId,
						EntryNamespaceId:   cne.EntryNamespaceId,
						Path:               strings.Join(append(teamFolderEntry.ParentIds, ""), "/"),
						PathDisplay:        strings.Join(append(teamFolderEntry.ParentNames, cne.Name), "/"),
						SummaryMemberCount: SummaryMemberCount{},
						SummaryEntryCount:  SummaryEntryCount{},
						Size:               cne.Size,
					})
					continue
				}

				childEntries = append(childEntries, &SummarizeTeamFolderEntry{
					TeamFolder:  teamFolderEntry.TeamFolder,
					FileId:      cne.FileId,
					ParentId:    ne.FileId,
					ParentIds:   append(teamFolderEntry.ParentIds, ne.FileId),
					ParentNames: append(teamFolderEntry.ParentNames, ne.Name),
				})
			}
		case errors.Is(err, gorm.ErrRecordNotFound):
			l.Debug("No child entries", esl.Error(err))

		default:
			l.Debug("Unable to find namespace entry", esl.Error(err))
			return err
		}
	}

	if len(childEntries) > 0 {
		q := s.Get(teamSummarizeTeamFolderEntry)
		for _, ce := range childEntries {
			q.Enqueue(ce)
		}
	}

	return nil
}
