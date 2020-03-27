package content

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"go.uber.org/zap"
)

type Member struct {
	Peer           rc_conn.ConnBusinessFile
	Metadata       kv_storage.Storage
	MembershipList kv_storage.Storage
	Tree           kv_storage.Storage
	Membership     rp_model.RowReport
	NoMember       rp_model.RowReport
}

type Membership struct {
	NamespaceId   string `json:"namespace_id"`
	NamespaceName string `json:"namespace_name"`
	Path          string `json:"path"`
	IsTeamFolder  bool   `json:"is_team_folder"`
	OwnerTeamId   string `json:"owner_team_id"`
	OwnerTeamName string `json:"owner_team_name"`
	AccessType    string `json:"access_type"`
	MemberType    string `json:"member_type"`
	MemberId      string `json:"member_id"`
	MemberName    string `json:"member_name"`
	MemberEmail   string `json:"member_email"`
}

type NoMember struct {
	OwnerTeamId   string `json:"owner_team_id"`
	OwnerTeamName string `json:"owner_team_name"`
	NamespaceId   string `json:"namespace_id"`
	NamespaceName string `json:"namespace_name"`
	Path          string `json:"path"`
}

func (z *Member) Preset() {
	z.Membership.SetModel(
		&Membership{},
		rp_model.HiddenColumns(
			"owner_team_id",
			"namespace_id",
			"namespace_name",
			"member_id",
		),
	)
	z.NoMember.SetModel(
		&NoMember{},
		rp_model.HiddenColumns(
			"owner_team_id",
			"namespace_id",
			"namespace_name",
		),
	)
}

func (z *Member) Exec(c app_control.Control) error {
	l := c.Log()

	q := c.NewQueue()
	s := &TeamScanner{
		ctx:   z.Peer.Context(),
		ctl:   c,
		queue: q,
		scanner: &ScanNamespaceMetadataAndMembership{
			metadata: &ScanNamespaceMetadata{
				metadata: z.Metadata,
				queue:    q,
			},
			membership: &ScanNamespaceMembership{
				membership: z.MembershipList,
				queue:      q,
			},
		},
	}
	if err := s.Scan(); err != nil {
		return err
	}
	q.Wait()

	if err := z.Membership.Open(); err != nil {
		return err
	}
	if err := z.NoMember.Open(); err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(z.Peer.Context()).Admin()
	if err != nil {
		return err
	}

	st := &TeamFolderScanner{
		ctl:      c,
		ctx:      z.Peer.Context().AsAdminId(admin.TeamMemberId),
		metadata: z.Metadata,
		tree:     z.Tree,
	}
	if err := st.Scan(); err != nil {
		return err
	}

	return z.Tree.View(func(treeKvs kv_kvs.Kvs) error {
		return treeKvs.ForEachModel(&Tree{}, func(key string, m interface{}) error {
			t := m.(*Tree)
			ll := l.With(zap.String("nsid", t.NamespaceId))
			ll.Debug("Preparing for report")
			meta := &mo_sharedfolder.SharedFolder{}
			err := z.Metadata.View(func(metaKvs kv_kvs.Kvs) error {
				return metaKvs.GetJsonModel(t.NamespaceId, meta)
			})
			if err != nil {
				ll.Debug("Unable to get metadata for the namespace", zap.Error(err))
				return err
			}
			return z.MembershipList.View(func(memberKvs kv_kvs.Kvs) error {
				members := make([]mo_sharedfolder_member.Metadata, 0)
				if err := memberKvs.GetJsonModel(t.NamespaceId, &members); err != nil {
					l.Debug("Unable to retrieve model", zap.Error(err))
					return err
				}
				if len(members) < 1 {
					z.NoMember.Row(&NoMember{
						OwnerTeamId:   meta.OwnerTeamId,
						OwnerTeamName: meta.OwnerTeamName,
						NamespaceId:   meta.SharedFolderId,
						NamespaceName: meta.Name,
						Path:          t.RelativePath,
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

					ms := &Membership{
						OwnerTeamId:   meta.OwnerTeamId,
						OwnerTeamName: meta.OwnerTeamName,
						NamespaceId:   meta.SharedFolderId,
						NamespaceName: meta.Name,
						Path:          t.RelativePath,
						IsTeamFolder:  meta.IsTeamFolder || meta.IsInsideTeamFolder,
						AccessType:    member.AccessType(),
						MemberType:    member.MemberType(),
						MemberId:      memberId,
						MemberName:    memberName,
						MemberEmail:   memberEmail,
					}

					z.Membership.Row(ms)
				}

				return nil
			})
		})
	})
}

func (z *Member) Test(c app_control.Control) error {
	return qt_errors.ErrorImplementMe
}
