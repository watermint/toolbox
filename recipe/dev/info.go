package dev

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Info struct {
	rc_recipe.RemarkSecret
}

func (z *Info) Preset() {
}

func (z *Info) Exec(c app_control.Control) error {
	return nil
}

func (z *Info) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Info{}, rc_recipe.NoCustomValues)
}
