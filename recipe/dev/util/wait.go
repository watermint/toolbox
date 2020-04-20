package util

import (
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"go.uber.org/zap"
	"time"
)

type Wait struct {
	rc_recipe.RemarkSecret
	Seconds mo_int.RangeInt
}

func (z *Wait) Exec(c app_control.Control) error {
	c.Log().Info("Wait", zap.Int("seconds", z.Seconds.Value()))
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
