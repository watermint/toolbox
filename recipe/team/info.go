package team

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type Info struct {
	Peer dbx_conn.ConnBusinessInfo
	Info rp_model.RowReport
}

func (z *Info) Preset() {
	z.Info.SetModel(&mo_team.Info{})
}

func (z *Info) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &Info{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "info", func(cols map[string]string) error {
		if _, ok := cols["team_id"]; !ok {
			return errors.New("`team_id` is not found")
		}
		return nil
	})
}

func (z *Info) Exec(c app_control.Control) error {
	if err := z.Info.Open(); err != nil {
		return err
	}

	info, err := sv_team.New(z.Peer.Context()).Info()
	if err != nil {
		return err
	}
	z.Info.Row(info)

	return nil
}
