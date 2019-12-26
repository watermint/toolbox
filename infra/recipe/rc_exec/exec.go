package rc_exec

import (
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"go.uber.org/zap"
)

func Exec(ctl app_control.Control, r rc_recipe.Recipe, custom func(r rc_recipe.Recipe)) error {
	return ExecSpec(ctl, rc_spec.New(r), custom)
}

func ExecSpec(ctl app_control.Control, spec rc_recipe.Spec, custom func(r rc_recipe.Recipe)) error {
	l := ctl.Log()
	if spec == nil {
		l.Debug("Spec not found")
		return errors.New("no spec found")
	}
	l = l.With(zap.String("cliPath", spec.CliPath()))
	scr, _, err := spec.ApplyValues(ctl, custom)
	if err != nil {
		return err
	}
	defer func() {
		err = spec.SpinDown(ctl)
		l.Debug("Spin down failed", zap.Error(err))
	}()
	return scr.Exec(rc_kitchen.NewKitchen(ctl, scr))
}
