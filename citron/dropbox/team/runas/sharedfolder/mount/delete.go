package mount

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Delete struct {
	Peer                 dbx_conn.ConnScopedTeam
	SharedFolderId       string
	MemberEmail          string
	Mount                rp_model.RowReport
	InfoAlreadyUnmounted app_msg.Message
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
	)
	z.Mount.SetModel(
		&mo_sharedfolder.SharedFolder{},
		rp_model.HiddenColumns(
			"parent_shared_folder_id",
			"team_member_id",
			"namespace_id",
			"owner_team_id",
		),
	)
}

func (z *Delete) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Mount.Open(); err != nil {
		return err
	}

	member, err := sv_member.New(z.Peer.Client()).ResolveByEmail(z.MemberEmail)
	if err != nil {
		return err
	}

	ctx := z.Peer.Client().AsMemberId(member.TeamMemberId).WithPath(dbx_client.Namespace(member.Profile().RootNamespaceId))

	l.Debug("Member found", esl.Any("member", member))

	err = sv_sharedfolder_mount.New(ctx).Unmount(&mo_sharedfolder.SharedFolder{SharedFolderId: z.SharedFolderId})
	if err != nil {
		de := dbx_error.NewErrors(err)
		switch {
		case de.HasPrefix("access_error/unmounted"):
			c.UI().Info(z.InfoAlreadyUnmounted)

		default:
			return err
		}
	}

	mount, err := sv_sharedfolder.New(ctx).Resolve(z.SharedFolderId)
	if err != nil {
		return err
	}
	z.Mount.Row(mount)

	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
