package content

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
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
	Membership     rp_model.RowReport
}

type Membership struct {
	OwnerTeamId   string `json:"owner_team_id"`
	OwnerTeamName string `json:"owner_team_name"`
	NamespaceId   string `json:"namespace_id"`
	NamespaceName string `json:"namespace_name"`
	Path          string `json:"path"`
	AccessType    string `json:"access_type"`
	MemberType    string `json:"member_type"`
	MemberId      string `json:"member_id"`
	MemberName    string `json:"member_name"`
	MemberEmail   string `json:"member_email"`
}

func (z *Member) Preset() {
	z.Membership.SetModel(&Membership{})
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
	if err := s.NamespacesOfTeam(); err != nil {
		return err
	}
	if err := s.NamespacesOfMembers(); err != nil {
		return err
	}
	q.Wait()

	if err := z.Membership.Open(); err != nil {
		return err
	}

	return z.Metadata.View(func(metaKvs kv_kvs.Kvs) error {
		return metaKvs.ForEachModel(&mo_sharedfolder.SharedFolder{}, func(key string, m interface{}) error {
			meta := m.(*mo_sharedfolder.SharedFolder)
			l.Debug("Namespace", zap.Any("metadata", meta))

			if meta.IsInsideTeamFolder || meta.IsTeamFolder {
				l.Debug("Team folder or inside team folder")
			}

			return z.MembershipList.View(func(membershipKvs kv_kvs.Kvs) error {
				members := make([]mo_sharedfolder_member.Metadata, 0)
				if err := membershipKvs.GetJsonModel(key, &members); err != nil {
					l.Debug("Unable to retrieve model", zap.Error(err))
					return err
				}

				for _, member := range members {
					ms := &Membership{
						OwnerTeamId:   meta.OwnerTeamId,
						OwnerTeamName: meta.OwnerTeamName,
						NamespaceId:   meta.SharedFolderId,
						NamespaceName: meta.Name,
						Path:          meta.PathLower,
						AccessType:    member.AccessType(),
						MemberType:    member.MemberType(),
						MemberId:      "",
						MemberName:    "",
						MemberEmail:   "",
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
