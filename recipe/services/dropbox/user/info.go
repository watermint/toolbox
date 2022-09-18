package user

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Info struct {
	Peer    dbx_conn.ConnScopedIndividual
	Profile rp_model.RowReport
}

func (z *Info) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeAccountInfoRead,
	)
	z.Profile.SetModel(
		&mo_member.Member{},
		rp_model.HiddenColumns(
			"team_member_id",
			"external_id",
			"account_id",
			"persistent_id",
			"member_folder_id",
			"abbreviated_name",
			"familiar_name",
			"role",
			"tag",
		),
	)
}

func (z *Info) Exec(c app_control.Control) error {
	if err := z.Profile.Open(); err != nil {
		return err
	}
	profile, err := sv_profile.NewProfile(z.Peer.Client()).Current()
	if err != nil {
		return err
	}
	z.Profile.Row(profile)
	return nil
}

func (z *Info) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Info{}, rc_recipe.NoCustomValues)
}
