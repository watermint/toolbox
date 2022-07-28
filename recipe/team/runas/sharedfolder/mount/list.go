package mount

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type List struct {
	Peer        dbx_conn.ConnScopedTeam
	MemberEmail string
	Mounts      rp_model.RowReport
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeTeamDataTeamSpace,
	)
	z.Mounts.SetModel(
		&mo_sharedfolder.SharedFolder{},
		rp_model.HiddenColumns(
			"parent_shared_folder_id",
			"team_member_id",
			"namespace_id",
			"owner_team_id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Mounts.Open(); err != nil {
		return err
	}

	member, err := sv_member.New(z.Peer.Context()).ResolveByEmail(z.MemberEmail)
	if err != nil {
		return err
	}

	l.Debug("Member found", esl.Any("member", member))

	ctx := z.Peer.Context().AsMemberId(member.TeamMemberId)
	mounts, err := sv_sharedfolder_mount.New(ctx).List()
	if err != nil {
		return err
	}

	for _, mount := range mounts {
		z.Mounts.Row(mount)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
