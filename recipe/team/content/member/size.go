package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_folder_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_member_folder"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_content"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_teamfolder_scanner"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/multierr"
)

type Size struct {
	Peer                           dbx_conn.ConnScopedTeam
	FolderMember                   kv_storage.Storage
	FolderOrphaned                 kv_storage.Storage
	Folder                         mo_filter.Filter
	MemberType                     mo_filter.Filter
	ScanTimeout                    mo_string.SelectString
	MemberCount                    rp_model.RowReport
	IncludeSubFolders              bool
	ErrorUnableToScanMemberFolders app_msg.Message
}

func (z *Size) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.Folder.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
	z.ScanTimeout.SetOptions(string(uc_teamfolder_scanner.ScanTimeoutShort),
		string(uc_teamfolder_scanner.ScanTimeoutShort),
		string(uc_teamfolder_scanner.ScanTimeoutLong),
	)
	z.MemberCount.SetModel(
		&uc_team_content.MemberCount{},
		rp_model.HiddenColumns(
			"namespace_id",
			"namespace_name",
			"parent_namespace_id",
			"owner_team_id",
			"member_emails",
		),
	)
}

func (z *Size) Exec(c app_control.Control) error {
	l := c.Log()

	if err := z.MemberCount.Open(); err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(z.Peer.Context()).Admin()
	if err != nil {
		return err
	}
	team, err := sv_team.New(z.Peer.Context()).Info()
	if err != nil {
		return err
	}

	teamFolderScanner := uc_teamfolder_scanner.New(c, z.Peer.Context(), uc_teamfolder_scanner.ScanTimeoutMode(z.ScanTimeout.Value()))
	teamFolders, err := teamFolderScanner.Scan(z.Folder)
	if err != nil {
		return err
	}

	memberFolderScanner := uc_member_folder.New(c, z.Peer.Context())
	memberFolders, err := memberFolderScanner.Scan(z.Folder)
	if err != nil {
		l.Debug("Failed to scan member folders", esl.Error(err))
		c.UI().Error(z.ErrorUnableToScanMemberFolders.With("Error", err))
		memberFolders = make([]*uc_member_folder.MemberNamespace, 0)
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_folder_members", uc_folder_member.ScanFolderMember, z.Peer.Context(), z.FolderMember, z.FolderOrphaned)
		q := s.Get("scan_folder_members")

		for _, tf := range teamFolders {
			q.Enqueue(&uc_folder_member.FolderEntry{
				Folder:    tf.TeamFolder,
				Path:      "",
				AsAdminId: admin.TeamMemberId,
			})

			for p, nf := range tf.NestedFolders {
				q.Enqueue(&uc_folder_member.FolderEntry{
					Folder:    nf,
					Path:      p,
					AsAdminId: admin.TeamMemberId,
				})
			}
		}

		for _, mf := range memberFolders {
			if (mf.Namespace.IsTeamFolder || mf.Namespace.IsInsideTeamFolder) && mf.Namespace.OwnerTeamId == team.TeamId {
				l.Debug("Skip team folder", esl.Any("folder", mf))
				continue
			}
			q.Enqueue(&uc_folder_member.FolderEntry{
				Folder:     mf.Namespace,
				AsMemberId: mf.TeamMemberId,
			})
		}
	})

	svg := sv_group_member.NewCachedDirectory(z.Peer.Context())
	counts := make(map[string]uc_team_content.MemberCount)

	err0 := z.FolderMember.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&uc_folder_member.FolderMemberEntry{}, func(key string, m interface{}) error {
			entry := m.(*uc_folder_member.FolderMemberEntry)
			mc, err := uc_team_content.MemberCountFromMemberEntry(entry.Path, entry.Folder, entry.Member, svg)
			if err != nil {
				l.Debug("Unable to resolve member count", esl.Error(err))
				return err
			}
			if c, ok := counts[entry.Folder.SharedFolderId]; ok {
				counts[entry.Folder.SharedFolderId] = c.Merge(mc)
			} else {
				counts[entry.Folder.SharedFolderId] = mc
			}
			return nil
		})
	})
	err1 := z.FolderOrphaned.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&uc_folder_member.FolderNoMemberEntry{}, func(key string, m interface{}) error {
			entry := m.(*uc_folder_member.FolderNoMemberEntry)
			l.Debug("folderMember", esl.Any("entry", entry))
			mc := uc_team_content.MemberCountFromSharedFolder(entry.Path, entry.Folder)
			if c, ok := counts[entry.Folder.SharedFolderId]; ok {
				counts[entry.Folder.SharedFolderId] = c.Merge(mc)
			} else {
				counts[entry.Folder.SharedFolderId] = mc
			}
			return nil
		})
	})

	for _, mc := range counts {
		if c, ok := counts[mc.ParentNamespaceId]; ok {
			counts[mc.ParentNamespaceId] = c.Merge(mc)
		}
	}

	for _, mc := range counts {
		if z.IncludeSubFolders || mc.ParentNamespaceId == "" {
			z.MemberCount.Row(mc)
		}
	}

	return multierr.Combine(err0, err1)
}

func (z *Size) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Size{}, rc_recipe.NoCustomValues)
}
