package rc_exec

import (
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"go.uber.org/zap"
)

func Exec(ctl app_control.Control, r rc_recipe.Recipe, custom func(r rc_recipe.Recipe)) error {
	return ExecSpec(ctl, rc_spec.New(r), custom)
}

func ExecSpec(ctl app_control.Control, spec rc_recipe.Spec, custom func(r rc_recipe.Recipe)) error {
	return DoSpec(ctl, spec, custom, func(r rc_recipe.Recipe, ctl app_control.Control) error {
		return r.Exec(ctl)
	})
}

func DoSpec(ctl app_control.Control, spec rc_recipe.Spec, custom func(r rc_recipe.Recipe), do func(r rc_recipe.Recipe, ctl app_control.Control) error) error {
	l := ctl.Log()
	if spec == nil {
		l.Debug("Spec not found")
		return errors.New("no spec found")
	}
	l = l.With(zap.String("cliPath", spec.CliPath()))
	scr, err := spec.SpinUp(ctl, custom)
	if err != nil {
		return err
	}

	rcpErr := do(scr, ctl)
	if err = spec.SpinDown(ctl); err != nil {
		l.Debug("Spin down error", zap.Error(err))
	}
	return rcpErr
}
