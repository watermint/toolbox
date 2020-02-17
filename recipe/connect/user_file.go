package connect

import (
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type UserFile struct {
	Peer    rc_conn.ConnUserFile
	Success app_msg.Message
	Failure app_msg.Message
}

func (z *UserFile) Preset() {
}

func (z *UserFile) Exec(c app_control.Control) error {
	ui := c.UI()
	user, err := sv_profile.NewProfile(z.Peer.Context()).Current()
	if err != nil {
		ui.Failure(z.Failure.With("Error", err))
		return err
	}
	ui.Success(z.Success.With("UserEmail", user.Email))
	return nil
}

func (z *UserFile) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}
