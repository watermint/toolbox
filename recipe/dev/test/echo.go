package test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Echo struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkTransient
	Text string
}

func (z *Echo) Exec(c app_control.Control) error {
	c.UI().Info(app_msg.Raw(z.Text))
	return nil
}

func (z *Echo) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Echo{}, func(r rc_recipe.Recipe) {
		m := r.(*Echo)
		m.Text = "Hello, World"
	})
}

func (z *Echo) Preset() {
}
