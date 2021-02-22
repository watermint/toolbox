package date

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"time"
)

type Today struct {
	rc_recipe.RemarkTransient
	Utc bool
}

func (z *Today) Preset() {
}

func (z *Today) Exec(c app_control.Control) error {
	if z.Utc {
		ui_out.TextOut(c, time.Now().UTC().Format("2006-01-02"))
	} else {
		ui_out.TextOut(c, time.Now().Format("2006-01-02"))
	}
	return nil
}

func (z *Today) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Today{}, rc_recipe.NoCustomValues)
}
