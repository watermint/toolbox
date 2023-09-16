package uc_insight

import "github.com/watermint/toolbox/essentials/log/esl"

type SummaryEntry struct {
	// primary keys
	FileId string `path:"file_id" gorm:"primaryKey"`

	// attributes
	Name             string `path:"name"`
	EntryType        string `path:"entry_type"`
	ParentFolderId   string `path:"parent_folder_id" gorm:"index"`
	EntryNamespaceId string `path:"entry_namespace_id" gorm:"index"`

	// InheritType is the type of inheritance.
	// "no_inherit" means the entry is not inherited from the parent folder.
	// "inherit" means the entry is inherited from the parent folder.
	// "inherit_plus" means the entry is inherited from the parent folder and have additional permissions.
	InheritType        string `path:"inherit_type"`
	CountUnique        uint64 `path:"count_unique"`
	CountMember        uint64 `path:"count_member"`
	CountInvitee       uint64 `path:"count_invitee"`
	CountGroup         uint64 `path:"count_group"`
	CountGroupExternal uint64 `path:"count_group_external"`

	// Links
	CountLinks uint64 `path:"count_links"`

	// Items
	CountEntries   uint64 `path:"count_entries"`
	CountFiles     uint64 `path:"count_files"`
	CountFolders   uint64 `path:"count_folders"`
	CountNamespace uint64 `path:"count_namespace"`

	// Size
	Size uint64 `path:"size"`

	// Updated is the timestamp when the entry is updated.
	Updated uint64 `gorm:"autoUpdateTime"`
}

func (z SummaryEntry) AddAccess(am AccessMember) SummaryEntry {
	switch am.MemberType {
	case "user":
		z.CountMember++
	case "group":
		z.CountGroup++
		if am.SameTeam == "no" {
			z.CountGroupExternal++
		}
	case "invitee":
		z.CountInvitee++
	}

	return z
}

func (z tsImpl) summarizeEntry(fileId string) error {
	l := z.ctl.Log().With(esl.String("fileId", fileId))
	entry := &SummaryEntry{}

	ne := &NamespaceEntry{}
	if err := z.db.First(ne, "file_id = ?", fileId).Error; err != nil {
		l.Debug("cannot find entry", esl.Error(err))
		return err
	}

	entry.FileId = fileId
	entry.Name = ne.Name
	entry.EntryType = ne.EntryType
	entry.ParentFolderId = ne.ParentFolderId
	entry.EntryNamespaceId = ne.EntryNamespaceId

	var linkCount int64
	z.db.Model(&SharedLink{}).Where("file_id = ?", fileId).Count(&linkCount)
	entry.CountLinks = uint64(linkCount)

	switch ne.EntryType {
	case "file":
		fm := &FileMember{}
		row, err := z.db.Model(fm).Where("file_id = ?", fileId).Rows()
		if err != nil {
			l.Debug("cannot find file members", esl.Error(err))
			return err
		}
		defer func() {
			_ = row.Close()
		}()

		for row.Next() {
			if err := z.db.ScanRows(row, fm); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				return err
			}
			entry.InheritType = "inherit_plus"

		}
	}

	z.db.Save(entry)

	return nil
}
