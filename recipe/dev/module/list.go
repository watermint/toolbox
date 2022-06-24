package module

import (
	"github.com/watermint/toolbox/essentials/go/go_module"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type List struct {
	rc_recipe.RemarkSecret
}

func (z *List) Preset() {
}

func (z *List) Exec(c app_control.Control) error {
	b, err := go_module.ScanBuild()
	if err != nil {
		return err
	}
	l := c.Log()
	l.Info("Go Version", esl.String("version", b.GoVersion()))
	for _, m := range b.Modules() {
		l.Info("Module", esl.String("Path", m.Path()), esl.String("Version", m.Version()))
	}

	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}
