package account

import (
	"github.com/watermint/toolbox/domain/figma/api/fg_conn"
	"github.com/watermint/toolbox/domain/figma/model/mo_user"
	"github.com/watermint/toolbox/domain/figma/service/sv_user"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Info struct {
	Peer fg_conn.ConnFigmaApi
	User rp_model.RowReport
}

func (z *Info) Preset() {
	z.User.SetModel(&mo_user.User{})
}

func (z *Info) Exec(c app_control.Control) error {
	if err := z.User.Open(); err != nil {
		return err
	}
	user, err := sv_user.New(z.Peer.Client()).Current()
	if err != nil {
		return err
	}
	z.User.Row(user)
	return nil
}

func (z *Info) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Info{}, rc_recipe.NoCustomValues)
}
