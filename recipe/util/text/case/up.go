package _case

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"strings"
)

type Up struct {
	rc_recipe.RemarkTransient
	Text string
}

func (z *Up) Preset() {
}

func (z *Up) Exec(c app_control.Control) error {
	ui_out.TextOut(c, strings.ToUpper(z.Text))
	return nil
}

func (z *Up) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Up{}, func(r rc_recipe.Recipe) {
		m := r.(*Up)
		m.Text = "hello world"
	})
}
