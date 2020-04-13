package bootstrap

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Bootstrap struct {
}

func (z *Bootstrap) Preset() {
}

func (z *Bootstrap) Exec(c app_control.Control) error {
	if err := rc_exec.Exec(c, &Autodelete{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return nil
}

func (z *Bootstrap) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Bootstrap{}, rc_recipe.NoCustomValues)
}
