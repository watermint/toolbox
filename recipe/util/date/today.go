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
	Utc    bool
	Offset int
}

func (z *Today) Preset() {
}

func (z *Today) Exec(c app_control.Control) error {
	t := time.Now().Add(time.Duration(z.Offset) * time.Hour * 24)
	if z.Utc {
		ui_out.TextOut(c, t.UTC().Format("2006-01-02"))
	} else {
		ui_out.TextOut(c, t.Format("2006-01-02"))
	}
	return nil
}

func (z *Today) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Today{}, rc_recipe.NoCustomValues)
}
