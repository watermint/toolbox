package bootstrap

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/job"
)

type OptInFeatureAutodelete struct {
	app_feature.OptInStatus
}

type Autodelete struct {
	Days int
}

func (z *Autodelete) Preset() {
	z.Days = 7
}

func (z *Autodelete) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()
	if f, found := c.Feature().OptInGet(&OptInFeatureAutodelete{}); !found {
		l.Debug("Skip cleanup")
		return nil
	} else if !f.OptInIsEnabled() {
		l.Debug("The feature disabled")
		return nil
	} else {
		ui.Info(app_feature.OptInDisclaimer(f))
	}
	return rc_exec.Exec(c, &job.Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*job.Delete)
		m.Days = z.Days
	})
}

func (z *Autodelete) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Autodelete{}, rc_recipe.NoCustomValues)
}
