package mount

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Add struct {
	Peer               dbx_conn.ConnScopedTeam
	MemberEmail        string
	SharedFolderId     string
	Mount              rp_model.RowReport
	InfoAlreadyMounted app_msg.Message
	BasePath           mo_string.SelectString
}

func (z *Add) Preset() {
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
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Add) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Mount.Open(); err != nil {
		return err
	}

	member, err := sv_member.New(z.Peer.Client()).ResolveByEmail(z.MemberEmail)
	if err != nil {
		return err
	}

	l.Debug("Member found", esl.Any("member", member))

	mount, err := sv_sharedfolder_mount.New(
		z.Peer.Client().
			AsMemberId(member.TeamMemberId, dbx_filesystem.AsNamespaceType(z.BasePath.Value())).
			WithPath(dbx_client.Namespace(member.Profile().RootNamespaceId))).
		Mount(&mo_sharedfolder.SharedFolder{SharedFolderId: z.SharedFolderId})
	if err != nil {
		de := dbx_error.NewErrors(err)
		switch {
		case de.HasPrefix("already_mounted"):
			c.UI().Info(z.InfoAlreadyMounted)

		default:
			return err
		}
	}

	z.Mount.Row(mount)
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
