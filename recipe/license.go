package recipe

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_license"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type License struct {
	rc_recipe.RemarkTransient
}

func (z *License) Preset() {
}

func (z *License) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}

func (z *License) Exec(c app_control.Control) error {
	return dc_license.Generate(c, c.UI())
}
