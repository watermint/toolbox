package util

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"time"
)

type Wait struct {
	rc_recipe.RemarkConsole
	rc_recipe.RemarkSecret
	Seconds mo_int.RangeInt
}

func (z *Wait) Exec(c app_control.Control) error {
	c.Log().Info("Wait", esl.Int("seconds", z.Seconds.Value()))
	time.Sleep(time.Duration(z.Seconds.Value()) * 1000 * time.Millisecond)
	return nil
}

func (z *Wait) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Wait{}, func(r rc_recipe.Recipe) {
		m := r.(*Wait)
		m.Seconds.SetValue(1)
	})
}

func (z *Wait) Preset() {
	z.Seconds.SetRange(1, 86400*7, 1)
}
