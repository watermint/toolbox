package uc_folder_member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_content"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"strings"
)

type FolderEntry struct {
	Folder     *mo_sharedfolder.SharedFolder `json:"folder"`
	Path       string                        `json:"path"`
	AsAdminId  string                        `json:"as_admin_id"`
	AsMemberId string                        `json:"as_member_id"`
}

type FolderMemberEntry struct {
	Folder *mo_sharedfolder.SharedFolder    `json:"folder"`
	Path   string                           `json:"path"`
	Member *mo_sharedfolder_member.Metadata `json:"member"`
}

func (z FolderMemberEntry) ToMembership() *uc_team_content.Membership {
	return uc_team_content.NewMembership(z.Folder, z.Path, z.Member)
}

type FolderNoMemberEntry struct {
	Folder *mo_sharedfolder.SharedFolder `json:"folder"`
	Path   string                        `json:"path"`
}

func (z FolderNoMemberEntry) ToNoMember() *uc_team_content.NoMember {
	return uc_team_content.NewNoMember(z.Folder, z.Path)
}

// Load folder member into storage (key -> *FolderMemberEntry)
func ScanFolderMember(entry *FolderEntry, ctx dbx_context.Context, storageFolderMember, storageNoMember kv_storage.Storage) error {
	l := ctx.Log().With(
		esl.String("NamespaceId", entry.Folder.SharedFolderId),
		esl.String("Path", entry.Path),
		esl.String("AsAdminId", entry.AsAdminId),
		esl.String("AsMemberId", entry.AsMemberId),
	)
	cta := ctx
	switch {
	case entry.AsMemberId != "":
		cta = ctx.AsMemberId(entry.AsMemberId)
	case entry.AsAdminId != "":
		cta = ctx.AsAdminId(entry.AsAdminId)
	}

	l.Debug("Scan folder members")
	members, err := sv_sharedfolder_member.NewBySharedFolderId(cta, entry.Folder.SharedFolderId).List()
	if err != nil {
		l.Debug("Unable to retrieve members", esl.Error(err))
		return err
	}

	if len(members) < 1 {
		l.Debug("No member found")
		return storageNoMember.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutJsonModel(entry.Folder.SharedFolderId, &FolderNoMemberEntry{
				Folder: entry.Folder,
				Path:   entry.Path,
			})
		})
	}

	l.Debug("Members found", esl.Int("memberCount", len(members)))

	for _, member := range members {
		memberKey := ""
		if u, ok := member.User(); ok {
			memberKey = u.AccountId
		}
		if g, ok := member.Group(); ok {
			memberKey = g.GroupId
		}
		if e, ok := member.Invitee(); ok {
			memberKey = e.InviteeEmail
		}

		key := strings.Join([]string{
			entry.Folder.SharedFolderId,
			entry.Path,
			memberKey,
		}, "/")

		err = storageFolderMember.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutJsonModel(key, &FolderMemberEntry{
				Folder: entry.Folder,
				Path:   entry.Path,
				Member: member.Metadata(),
			})
		})
		if err != nil {
			return err
		}
	}
	return nil
}
