package team

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_team"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type Feature struct {
	Peer    rc_conn.ConnBusinessInfo
	Feature rp_model.RowReport
}

func (z *Feature) Preset() {
	z.Feature.SetModel(&mo_team.Feature{})
}

func (z *Feature) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &Feature{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "feature", func(cols map[string]string) error {
		if _, ok := cols["upload_api_rate_limit"]; !ok {
			return errors.New("`upload_api_rate_limit` is not found")
		}
		return nil
	})
}

func (z *Feature) Exec(k rc_kitchen.Kitchen) error {
	if err := z.Feature.Open(); err != nil {
		return err
	}

	info, err := sv_team.New(z.Peer.Context()).Feature()
	if err != nil {
		return err
	}
	z.Feature.Row(info)

	return nil
}
