package util

import (
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"go.uber.org/zap"
	"time"
)

type Wait struct {
	Seconds int
}

func (z *Wait) Exec(c app_control.Control) error {
	if z.Seconds < 1 {
		return errors.New("seconds must grater than 1")
	}
	c.Log().Info("Wait", zap.Int("seconds", z.Seconds))
	time.Sleep(time.Duration(z.Seconds) * time.Second)
	return nil
}

func (z *Wait) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Wait{}, func(r rc_recipe.Recipe) {
		m := r.(*Wait)
		m.Seconds = 1
	})
}

func (z *Wait) Preset() {
	z.Seconds = 1
}
