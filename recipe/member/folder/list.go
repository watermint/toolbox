package folder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_folder_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_member_folder"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_content"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_teamfolder"
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
	Peer                           dbx_conn.ConnBusinessFile
	ScanTimeout                    mo_string.SelectString
	Member                         mo_filter.Filter
	Folder                         mo_filter.Filter
	FolderMember                   kv_storage.Storage
	FolderOrphaned                 kv_storage.Storage
	ErrorUnableToScanMemberFolders app_msg.Message
	MemberToFolder                 rp_model.RowReport
	MemberWithNoFolder             rp_model.RowReport
}

func (z *List) Preset() {
	z.Member.SetOptions(
		mo_filter.NewEmailFilter(),
	)
	z.Folder.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
	z.ScanTimeout.SetOptions(string(uc_teamfolder.ScanTimeoutShort),
		string(uc_teamfolder.ScanTimeoutShort),
		string(uc_teamfolder.ScanTimeoutLong),
	)
	z.MemberToFolder.SetModel(
		&uc_team_content.MemberToFolder{},
		rp_model.HiddenColumns(
			"team_member_id",
			"namespace_id",
			"owner_team_id",
		),
	)
	z.MemberWithNoFolder.SetModel(
		&mo_member.Member{},
		rp_model.HiddenColumns(
			"team_member_id",
			"email_verified",
			"familiar_name",
			"abbreviated_name",
			"member_folder_id",
			"external_id",
			"account_id",
			"persistent_id",
			"joined_on",
			"role",
			"tag",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()

	if err := z.MemberToFolder.Open(); err != nil {
		return err
	}
	if err := z.MemberWithNoFolder.Open(rp_model.NoConsoleOutput()); err != nil {
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
	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)
	memberInUse := make(map[string]bool)

	teamFolderScanner := uc_teamfolder.New(c, z.Peer.Context(), uc_teamfolder.ScanTimeoutMode(z.ScanTimeout.Value()))
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

	err1 := z.FolderMember.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&uc_folder_member.FolderMemberEntry{}, func(key string, m interface{}) error {
			entry := m.(*uc_folder_member.FolderMemberEntry)
			if m, ok := entry.Member.User(); ok {
				if !z.Member.Accept(m.Email) {
					l.Debug("Skip (email)", esl.Any("entry", entry))
					return nil
				}
				if !m.IsSameTeam {
					l.Debug("Skip (external)", esl.Any("entry", entry))
					return nil
				}
				l.Debug("Member", esl.Any("user", m), esl.Any("entry", entry))
				z.MemberToFolder.Row(uc_team_content.NewMemberToFolderByUser(m, entry.Folder, entry.Path, entry.Member))
				memberInUse[m.Email] = true
			} else if m, ok := entry.Member.Invitee(); ok {
				if !z.Member.Accept(m.InviteeEmail) {
					l.Debug("Skip (email)", esl.Any("entry", entry))
					return nil
				}
				if _, ok := emailToMember[m.InviteeEmail]; !ok {
					l.Debug("Skip (external)", esl.Any("entry", entry))
					return nil
				}
				l.Debug("Invitee", esl.Any("invitee", m), esl.Any("entry", entry))
				z.MemberToFolder.Row(uc_team_content.NewMemberToFolderByInvitee(m, entry.Folder, entry.Path, entry.Member))
				memberInUse[m.InviteeEmail] = true

			} else {
				l.Debug("Skip (member type)", esl.Any("entry", entry))
				return nil
			}

			return nil
		})
	})

	for _, member := range members {
		if t, ok := memberInUse[member.Email]; ok && t {
			l.Debug("Member is in use", esl.Any("member", member))
			continue
		}
		z.MemberWithNoFolder.Row(member)
	}
	return err1
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
