package folder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
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
)

type List struct {
	Peer                           dbx_conn.ConnScopedTeam
	ScanTimeout                    mo_string.SelectString
	Folder                         mo_filter.Filter
	Group                          mo_filter.Filter
	IncludeExternalGroups          bool
	FolderMember                   kv_storage.Storage
	FolderOrphaned                 kv_storage.Storage
	ErrorUnableToScanMemberFolders app_msg.Message
	GroupToFolder                  rp_model.RowReport
	GroupWithNoFolders             rp_model.RowReport
	BasePath                       mo_string.SelectString
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.Group.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
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
	z.GroupToFolder.SetModel(
		&uc_team_content.GroupToFolder{},
		rp_model.HiddenColumns(
			"namespace_id",
			"owner_team_id",
		),
	)
	z.GroupWithNoFolders.SetModel(
		&mo_group.Group{},
		rp_model.HiddenColumns(
			"group_id",
			"group_external_id",
		),
	)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()

	if err := z.GroupToFolder.Open(); err != nil {
		return err
	}
	if err := z.GroupWithNoFolders.Open(rp_model.NoConsoleOutput()); err != nil {
		return err
	}
	admin, err := sv_profile.NewTeam(z.Peer.Client()).Admin()
	if err != nil {
		return err
	}
	team, err := sv_team.New(z.Peer.Client()).Info()
	if err != nil {
		return err
	}
	groups, err := sv_group.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}
	groupInUse := make(map[string]bool)

	teamFolderScanner := uc_teamfolder_scanner.New(
		c,
		z.Peer.Client(),
		uc_teamfolder_scanner.ScanTimeoutMode(z.ScanTimeout.Value()),
		dbx_filesystem.AsNamespaceType(z.BasePath.Value()),
	)
	teamFolders, err := teamFolderScanner.Scan(z.Folder)
	if err != nil {
		return err
	}

	memberFolderScanner := uc_member_folder.New(c, z.Peer.Client(), dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	memberFolders, err := memberFolderScanner.Scan(z.Folder)
	if err != nil {
		l.Debug("Failed to scan member folders", esl.Error(err))
		c.UI().Error(z.ErrorUnableToScanMemberFolders.With("Error", err))
		memberFolders = make([]*uc_member_folder.MemberNamespace, 0)
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_folder_members",
			uc_folder_member.ScanFolderMember,
			z.Peer.Client(),
			z.FolderMember,
			z.FolderOrphaned,
			dbx_filesystem.AsNamespaceType(z.BasePath.Value()),
		)
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

	err1 := z.FolderMember.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&uc_folder_member.FolderMemberEntry{}, func(key string, m interface{}) error {
			entry := m.(*uc_folder_member.FolderMemberEntry)
			if g, ok := entry.Member.Group(); ok {
				if !z.Group.Accept(g.GroupName) {
					l.Debug("Skip (name)", esl.Any("entry", entry))
					return nil
				}
				if !z.IncludeExternalGroups && !g.IsSameTeam {
					l.Debug("Skip (external)", esl.Any("entry", entry))
					return nil
				}
				l.Debug("Group", esl.Any("group", g), esl.Any("entry", entry))
				z.GroupToFolder.Row(uc_team_content.NewGroupToFolder(g, entry.Folder, entry.Path, entry.Member))
				groupInUse[g.GroupId] = true

			} else {
				l.Debug("Skip (member type)", esl.Any("entry", entry))
				return nil
			}

			return nil
		})
	})

	for _, group := range groups {
		if t, ok := groupInUse[group.GroupId]; ok && t {
			l.Debug("Group is in use", esl.Any("group", group))
			continue
		}
		z.GroupWithNoFolders.Row(group)
	}
	return err1
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
