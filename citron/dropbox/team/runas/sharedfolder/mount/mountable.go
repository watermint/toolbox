package mount

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Mountable struct {
	Peer           dbx_conn.ConnScopedTeam
	Mountables     rp_model.RowReport
	MemberEmail    string
	IncludeMounted bool
	BasePath       mo_string.SelectString
}

func (z *Mountable) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
	)
	z.Mountables.SetModel(
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

func (z *Mountable) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Mountables.Open(); err != nil {
		return err
	}
	member, err := sv_member.New(z.Peer.Client()).ResolveByEmail(z.MemberEmail)
	if err != nil {
		return err
	}

	l.Debug("Member found", esl.Any("member", member))

	ctx := z.Peer.Client().
		AsMemberId(member.TeamMemberId, dbx_filesystem.AsNamespaceType(z.BasePath.Value())).
		WithPath(dbx_client.Namespace(member.Profile().RootNamespaceId))
	mounts, err := sv_sharedfolder_mount.New(ctx).List()
	if err != nil {
		return err
	}

	for _, mount := range mounts {
		if mount.PathLower != "" {
			if !z.IncludeMounted {
				continue
			}
		}
		z.Mountables.Row(mount)
	}
	return nil
}

func (z *Mountable) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
