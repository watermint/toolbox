package placeholder

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

// Prune is a placeholder recipe of keeping messages of prune related documentation.
type Prune struct {
	rc_recipe.RemarkSecret
}

func (z *Prune) Preset() {
}

func (z *Prune) Exec(c app_control.Control) error {
	return nil
}

func (z *Prune) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Prune{}, rc_recipe.NoCustomValues)
}
