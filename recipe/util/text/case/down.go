package _case

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"strings"
)

type Down struct {
	rc_recipe.RemarkTransient
	Text string
}

func (z *Down) Preset() {
}

func (z *Down) Exec(c app_control.Control) error {
	ui_out.TextOut(c, strings.ToLower(z.Text))
	return nil
}

func (z *Down) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Down{}, func(r rc_recipe.Recipe) {
		m := r.(*Down)
		m.Text = "hello world"
	})
}
