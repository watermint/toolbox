package group

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type List struct {
	Peer  rc_conn.ConnBusinessInfo
	Group rp_model.RowReport
}

func (z *List) Preset() {
	z.Group.SetModel(&mo_group.Group{})
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "group", func(cols map[string]string) error {
		if _, ok := cols["group_id"]; !ok {
			return errors.New("group_id is not found")
		}
		return nil
	})
}

func (z *List) Exec(c app_control.Control) error {
	groups, err := sv_group.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	if err := z.Group.Open(); err != nil {
		return err
	}
	for _, m := range groups {
		z.Group.Row(m)
	}
	return nil
}
