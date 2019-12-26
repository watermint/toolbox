package test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_runtime"
)

type Resources struct {
}

func (z *Resources) Preset() {
}

func (z *Resources) Test(c app_control.Control) error {
	return qt_endtoend.NoTestRequired()
}

func (z *Resources) Hidden() {
}

func (z *Resources) Exec(k rc_kitchen.Kitchen) error {
	qt_runtime.Suite(k.Control())
	return nil
}
