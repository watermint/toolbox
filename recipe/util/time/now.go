package time

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"time"
)

type Now struct {
	rc_recipe.RemarkTransient
	Utc bool
}

func (z *Now) Preset() {
}

func (z *Now) Exec(c app_control.Control) error {
	if z.Utc {
		ui_out.TextOut(c, time.Now().UTC().Format("15-04-05"))
	} else {
		ui_out.TextOut(c, time.Now().Format("15-04-05"))
	}
	return nil
}

func (z *Now) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Now{}, rc_recipe.NoCustomValues)
}
