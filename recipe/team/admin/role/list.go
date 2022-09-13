package role

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_adminrole"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_adminrole"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer  dbx_conn.ConnScopedTeam
	Roles rp_model.RowReport
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
	)
	z.Roles.SetModel(
		&mo_adminrole.Role{},
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Roles.Open(); err != nil {
		return err
	}

	roles, err := sv_adminrole.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	for _, role := range roles {
		z.Roles.Row(role)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
