package license

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_license_key"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Install struct {
	Key string
}

func (z *Install) Preset() {
}

func (z *Install) Exec(c app_control.Control) error {
	app_license_key.AddKey(c.Workspace(), z.Key)
	return nil
}

func (z *Install) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
