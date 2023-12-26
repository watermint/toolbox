package placeholder

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

// Pathchange is a placeholder recipe of keeping messages of path change related documentation.
type Pathchange struct {
	rc_recipe.RemarkSecret
}

func (z *Pathchange) Preset() {
}

func (z *Pathchange) Exec(c app_control.Control) error {
	return nil
}

func (z *Pathchange) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Pathchange{}, rc_recipe.NoCustomValues)
}
