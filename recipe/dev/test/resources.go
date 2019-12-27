package test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
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

func (z *Resources) Exec(c app_control.Control) error {
	qt_runtime.Suite(c)
	return nil
}
