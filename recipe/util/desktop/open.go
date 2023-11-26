package desktop

import (
	"github.com/watermint/toolbox/essentials/desktop/es_open"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Open struct {
	Path mo_path.FileSystemPath
}

func (z *Open) Preset() {
}

func (z *Open) Exec(c app_control.Control) error {
	return es_open.CurrentDesktop().Open(z.Path.Path()).Cause()
}

func (z *Open) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
