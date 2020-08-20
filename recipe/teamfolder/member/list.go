package member

import (
	"github.com/watermint/toolbox/domain/common/model/mo_filter"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_content"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer               dbx_conn.ConnBusinessFile
	Metadata           kv_storage.Storage
	MembershipList     kv_storage.Storage
	Tree               kv_storage.Storage
	Membership         rp_model.RowReport
	NoMember           rp_model.RowReport
	Folder             mo_filter.Filter
	MemberType         mo_filter.Filter
	memberTypeInternal mo_sharedfolder_member.FolderMemberFilter
	memberTypeExternal mo_sharedfolder_member.FolderMemberFilter
}

func (z *List) Preset() {
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
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()

	q := c.NewLegacyQueue()
	s := &uc_team_content.TeamScanner{
		Ctx:    z.Peer.Context(),
		Ctl:    c,
		Queue:  q,
		Filter: z.Folder,
		Scanner: &uc_team_content.ScanNamespaceMetadataAndMembership{
			Metadata: &uc_team_content.ScanNamespaceMetadata{
				Metadata: z.Metadata,
				Queue:    q,
			},
			Membership: &uc_team_content.ScanNamespaceMembership{
				Membership: z.MembershipList,
				Queue:      q,
			},
		},
	}
	if err := s.ScanTeamOnly(); err != nil {
		return err
	}
	q.Wait()

	if err := z.Membership.Open(); err != nil {
		return err
	}
	if err := z.NoMember.Open(); err != nil {
		return err
	}

	if z.memberTypeExternal.Enabled() || z.memberTypeInternal.Enabled() {
		members, err := sv_member.New(z.Peer.Context()).List()
		if err != nil {
			return err
		}
		z.memberTypeInternal.SetMembers(members)
		z.memberTypeExternal.SetMembers(members)
	}

	admin, err := sv_profile.NewTeam(z.Peer.Context()).Admin()
	if err != nil {
		return err
	}

	st := &uc_team_content.TeamFolderScanner{
		Ctl:      c,
		Ctx:      z.Peer.Context().AsAdminId(admin.TeamMemberId),
		Metadata: z.Metadata,
		Tree:     z.Tree,
	}
	if err := st.Scan(); err != nil {
		return err
	}

	return z.Tree.View(func(treeKvs kv_kvs.Kvs) error {
		return treeKvs.ForEachModel(&uc_team_content.Tree{}, func(key string, m interface{}) error {
			t := m.(*uc_team_content.Tree)
			ll := l.With(esl.String("nsid", t.NamespaceId))
			ll.Debug("Preparing for report")
			meta := &mo_sharedfolder.SharedFolder{}
			err := z.Metadata.View(func(metaKvs kv_kvs.Kvs) error {
				return metaKvs.GetJsonModel(t.NamespaceId, meta)
			})
			if err != nil {
				ll.Debug("Unable to get metadata for the namespace", esl.Error(err))
				return err
			}
			return z.MembershipList.View(func(memberKvs kv_kvs.Kvs) error {
				members := make([]mo_sharedfolder_member.Metadata, 0)
				if err := memberKvs.GetJsonModel(t.NamespaceId, &members); err != nil {
					l.Debug("Unable to retrieve model", esl.Error(err))
					return err
				}
				if len(members) < 1 {
					z.NoMember.Row(&uc_team_content.NoMember{
						OwnerTeamId:   meta.OwnerTeamId,
						OwnerTeamName: meta.OwnerTeamName,
						NamespaceId:   meta.SharedFolderId,
						NamespaceName: meta.Name,
						Path:          t.RelativePath,
						FolderType:    FolderType(meta),
					})
				}
				for _, member := range members {
					memberId := ""
					memberName := ""
					memberEmail := ""

					if u, ok := member.User(); ok {
						memberId = u.AccountId
						memberName = u.DisplayName
						memberEmail = u.Email
					}
					if g, ok := member.Group(); ok {
						memberId = g.GroupId
						memberName = g.GroupName
					}
					if e, ok := member.Invitee(); ok {
						memberEmail = e.InviteeEmail
					}

					ms := &uc_team_content.Membership{
						OwnerTeamId:   meta.OwnerTeamId,
						OwnerTeamName: meta.OwnerTeamName,
						NamespaceId:   meta.SharedFolderId,
						NamespaceName: meta.Name,
						Path:          t.RelativePath,
						FolderType:    FolderType(meta),
						AccessType:    member.AccessType(),
						MemberType:    member.MemberType(),
						MemberId:      memberId,
						MemberName:    memberName,
						MemberEmail:   memberEmail,
						SameTeam:      member.SameTeam(),
					}

					if z.MemberType.Accept(&member) {
						z.Membership.Row(ms)
					}
				}

				return nil
			})
		})
	})
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}

func FolderType(m *mo_sharedfolder.SharedFolder) string {
	switch {
	case m.IsTeamFolder, m.IsInsideTeamFolder:
		return "team_folder"
	default:
		return "shared_folder"
	}
}
