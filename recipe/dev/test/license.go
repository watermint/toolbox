package test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type License struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkLicenseRequired
}

func (z *License) Preset() {
}

func (z *License) Exec(c app_control.Control) error {
	c.Log().Info("License verification passed")
	return nil
}

func (z *License) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &License{}, rc_recipe.NoCustomValues)
}
