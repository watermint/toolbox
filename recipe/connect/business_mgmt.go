package connect

import (
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type BusinessMgmt struct {
	Peer    rc_conn.ConnBusinessMgmt
	Success app_msg.Message
	Failure app_msg.Message
}

func (z *BusinessMgmt) Preset() {
}

func (z *BusinessMgmt) Exec(c app_control.Control) error {
	ui := c.UI()
	admin, err := sv_profile.NewTeam(z.Peer.Context()).Admin()
	if err != nil {
		ui.Failure(z.Failure.With("Error", err))
		return err
	}
	ui.Success(z.Success.With("AdminEmail", admin.Email))
	return nil
}

func (z *BusinessMgmt) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}
