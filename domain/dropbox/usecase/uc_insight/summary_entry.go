package uc_insight

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"gorm.io/gorm"
)

type SummaryMemberCount struct {
	// CountUniqueLower is the lower bounds of number of unique members including users,
	// internal groups, and invitees.
	// But this count will not include number of members in external groups.
	// This count will not count difference between membership type like viewer or editor.
	CountUniqueLower uint64 `path:"count_unique_lower"`

	// CountUniqueUpper is the upper bounds of number of unique members including users,
	// is the lower bounds of number of unique members including users,
	// internal groups, and invitees.
	// But this count will not include number of members in external groups.
	// This count will not count difference between membership type like viewer or editor.
	CountUniqueUpper uint64 `path:"count_unique_upper"`

	// CountAccess is the number of unique direct access excluding groups (direct + invitees).
	// This count will not count difference between membership type like viewer or editor.
	CountMember uint64 `path:"count_member"`

	// CountMemberInternal is the number of unique internal members, this will not include invitees.
	// This count will not count difference between membership type like viewer or editor.
	CountMemberInternal uint64 `path:"count_member_internal"`

	// CountMemberExternal is the number of unique external members.
	// This count will not count difference between membership type like viewer or editor.
	CountMemberExternal uint64 `path:"count_member_external"`

	// CountInvitee is the number of unique invitees.
	// This count will not count difference between membership type like viewer or editor.
	CountInvitee uint64 `path:"count_invitee"`

	// CountGroup is the number of unique groups that includes both internal and external groups.
	// This count will not count difference between membership type like viewer or editor.
	CountGroup uint64 `path:"count_group"`

	// CountGroupExternal is the number of unique external groups.
	// This count will not count difference between membership type like viewer or editor.
	CountGroupExternal uint64 `path:"count_group_external"`
}

type SummaryEntryCount struct {

	// Links
	CountLinks uint64 `path:"count_links"`

	// Items
	CountEntries   uint64 `path:"count_entries"`
	CountFiles     uint64 `path:"count_files"`
	CountFolders   uint64 `path:"count_folders"`
	CountNamespace uint64 `path:"count_namespace"`
}

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
	InheritType string `path:"inherit_type"`
	SummaryMemberCount
	SummaryEntryCount

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

func (z summaryImpl) reduceMemberCount(accessMembers []*AccessMember) (smc SummaryMemberCount, err error) {
	l := z.ctl.Log().With(esl.Int("records", len(accessMembers)))
	directInternals := make(map[string]bool)
	directExternals := make(map[string]bool)
	invitees := make(map[string]bool)
	groupInternals := make(map[string]bool)
	groupExternals := make(map[string]bool)
	groupMemberExternals := make(map[string]int64)

	var summarizeInternalGroup func(am *AccessMember) error
	summarizeInternalGroup = func(am *AccessMember) error {
		groupInternals[am.GroupId] = true
		gm := &GroupMember{}
		rows, err := z.db.Model(&GroupMember{}).Where("group_id = ?", am.GroupId).Rows()
		if err != nil {
			l.Debug("cannot find group members", esl.Error(err))
			return err
		}
		defer func() {
			_ = rows.Close()
		}()
		for rows.Next() {
			gm = &GroupMember{}
			if err := z.db.ScanRows(rows, gm); err != nil {
				l.Debug("cannot scan row", esl.Error(err))
				return err
			}
			directInternals[am.UserAccountId] = true
		}
		return nil
	}

	for _, am := range accessMembers {
		switch am.MemberType {
		case "user":
			if am.SameTeam == "yes" {
				directInternals[am.UserAccountId] = true
			} else {
				directExternals[am.UserAccountId] = true
			}
		case "group":
			if am.SameTeam == "yes" {
				if err := summarizeInternalGroup(am); err != nil {
					l.Debug("cannot summarize internal group", esl.Error(err))
					return smc, err
				}

			} else {
				groupExternals[am.GroupId] = true
				ge := &Group{}
				if err := z.db.First(ge, "group_id = ?", am.GroupId).Error; err != nil {
					l.Debug("cannot find group", esl.Error(err))
					return smc, err
				}
				groupMemberExternals[am.GroupId] = int64(ge.MemberCount)
			}
		case "invitee":
			invitees[am.InviteeEmail] = true
		}
	}

	memberCountExternalLargestGroup := int64(0)
	memberCountExternalSum := int64(0)
	for _, count := range groupMemberExternals {
		memberCountExternalSum += count
		if count > memberCountExternalLargestGroup {
			memberCountExternalLargestGroup = count
		}
	}

	smc.CountMember = uint64(len(directInternals) + len(directExternals) + len(invitees))
	smc.CountMemberInternal = uint64(len(directInternals))
	smc.CountMemberExternal = uint64(len(directExternals))
	smc.CountUniqueLower = smc.CountMember + uint64(memberCountExternalLargestGroup)
	smc.CountUniqueUpper = smc.CountMember + uint64(memberCountExternalSum)
	smc.CountInvitee = uint64(len(invitees))
	smc.CountGroup = uint64(len(groupInternals) + len(groupExternals))
	smc.CountGroupExternal = uint64(len(groupExternals))
	return
}

func (z summaryImpl) summarizeEntry(fileId string) error {
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
	case "deleted", "file":
		return nil // ignore non folder entries

	case "folder":
		nm := &NamespaceMember{}
		row, err := z.db.Model(nm).Where("namespace_id = ?", ne.NamespaceId).Rows()
		switch {
		case err == nil:
			defer func() {
				_ = row.Close()
			}()

			acs := make([]*AccessMember, 0)
			for row.Next() {
				nm = &NamespaceMember{}
				if err := z.db.ScanRows(row, nm); err != nil {
					l.Debug("cannot scan row", esl.Error(err))
					return err
				}
				acs = append(acs, &nm.AccessMember)
			}
			entry.SummaryMemberCount, err = z.reduceMemberCount(acs)

			// Determine inherit type
			if ne.EntryNamespaceId == "" {
				entry.InheritType = "inherit"
			} else {
				nd := &NamespaceDetail{}
				if err := z.db.First(nd, "namespace_id = ?", ne.NamespaceId).Error; err != nil {
					entry.InheritType = ""
				} else {
					entry.InheritType = nd.AccessInheritance
				}
			}
			z.db.Save(entry)
			return nil

		case errors.Is(err, gorm.ErrRecordNotFound):
			entry.InheritType = "inherit"
			z.db.Save(entry)
			return nil

		default:
			l.Debug("cannot find namespace members", esl.Error(err))
			return err
		}

	default:
		l.Debug("unknown entry type", esl.String("entryType", ne.EntryType))
		return errors.New("unknown entry type")
	}
}
