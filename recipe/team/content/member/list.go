package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
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

type List struct {
	Peer                           dbx_conn.ConnScopedTeam
	FolderMember                   kv_storage.Storage
	FolderOrphaned                 kv_storage.Storage
	Membership                     rp_model.RowReport
	NoMember                       rp_model.RowReport
	Folder                         mo_filter.Filter
	MemberType                     mo_filter.Filter
	ScanTimeout                    mo_string.SelectString
	memberTypeInternal             mo_sharedfolder_member.FolderMemberFilter
	memberTypeExternal             mo_sharedfolder_member.FolderMemberFilter
	ErrorUnableToScanMemberFolders app_msg.Message
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.Membership.SetModel(
		&uc_team_content.Membership{},
		rp_model.HiddenColumns(
			"owner_team_id",
			"namespace_id",
			"namespace_name",
			"member_id",
		),
	)
	z.NoMember.SetModel(
		&uc_team_content.NoMember{},
		rp_model.HiddenColumns(
			"owner_team_id",
			"namespace_id",
			"namespace_name",
		),
	)
	z.Folder.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
	z.memberTypeInternal = mo_sharedfolder_member.NewInternalOpt()
	z.memberTypeExternal = mo_sharedfolder_member.NewExternalOpt()
	z.MemberType.SetOptions(
		z.memberTypeInternal,
		z.memberTypeExternal,
	)
	z.ScanTimeout.SetOptions(string(uc_teamfolder_scanner.ScanTimeoutShort),
		string(uc_teamfolder_scanner.ScanTimeoutShort),
		string(uc_teamfolder_scanner.ScanTimeoutLong),
	)
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()

	admin, err := sv_profile.NewTeam(z.Peer.Context()).Admin()
	if err != nil {
		return err
	}
	team, err := sv_team.New(z.Peer.Context()).Info()
	if err != nil {
		return err
	}

	if err := z.Membership.Open(); err != nil {
		return err
	}
	if err := z.NoMember.Open(rp_model.NoConsoleOutput()); err != nil {
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

	if z.memberTypeExternal.Enabled() || z.memberTypeInternal.Enabled() {
		members, err := sv_member.New(z.Peer.Context()).List()
		if err != nil {
			return err
		}
		z.memberTypeInternal.SetMembers(members)
		z.memberTypeExternal.SetMembers(members)
	}

	err0 := z.FolderMember.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&uc_folder_member.FolderMemberEntry{}, func(key string, m interface{}) error {
			entry := m.(*uc_folder_member.FolderMemberEntry)
			if z.MemberType.Accept(entry.Member) {
				z.Membership.Row(entry.ToMembership())
			}
			return nil
		})
	})
	err1 := z.FolderOrphaned.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&uc_folder_member.FolderNoMemberEntry{}, func(key string, m interface{}) error {
			entry := m.(*uc_folder_member.FolderNoMemberEntry)
			z.NoMember.Row(entry.ToNoMember())
			return nil
		})
	})
	return multierr.Combine(err0, err1)
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
