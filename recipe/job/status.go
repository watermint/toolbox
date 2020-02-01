package job

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

// Dummy command
type Status struct {
}

func (z *Status) Exec(c app_control.Control) error {
	return nil
}

func (z *Status) Test(c app_control.Control) error {
	return qt_endtoend.NoTestRequired()
}

func (z *Status) Preset() {
}
