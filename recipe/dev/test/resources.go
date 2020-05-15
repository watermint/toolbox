package test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_runtime"
)

type Resources struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkTransient
}

func (z *Resources) Preset() {
}

func (z *Resources) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}

func (z *Resources) Exec(c app_control.Control) error {
	qt_runtime.Suite(c)
	return nil
}
