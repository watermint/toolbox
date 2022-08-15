package datetime

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"time"
)

type Now struct {
	rc_recipe.RemarkTransient
	Utc        bool
	OffsetDay  int
	OffsetHour int
	OffsetMin  int
	OffsetSec  int
}

func (z *Now) Preset() {
}

func (z *Now) Exec(c app_control.Control) error {
	t := time.Now().
		Add(time.Duration(z.OffsetDay) * time.Hour * 24).
		Add(time.Duration(z.OffsetHour) * time.Hour).
		Add(time.Duration(z.OffsetMin) * time.Minute).
		Add(time.Duration(z.OffsetSec) * time.Second)

	if z.Utc {
		ui_out.TextOut(c, t.UTC().Format(time.RFC3339))
	} else {
		ui_out.TextOut(c, t.Format(time.RFC3339))
	}
	return nil
}

func (z *Now) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Now{}, rc_recipe.NoCustomValues)
}
