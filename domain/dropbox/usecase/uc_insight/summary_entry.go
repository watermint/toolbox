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
	CountAccess        uint64 `path:"count_access"`
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

	return nil
}
