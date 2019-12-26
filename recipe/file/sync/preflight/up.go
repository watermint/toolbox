package preflight

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/file"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type Up struct {
	Peer        rc_conn.ConnUserFile
	LocalPath   mo_path.FileSystemPath
	DropboxPath mo_path.DropboxPath
	Upload      *file.Upload
}

func (z *Up) Preset() {
}

func (z *Up) Console() {
}

func (z *Up) Exec(k rc_kitchen.Kitchen) error {
	return rc_exec.Exec(k.Control(), z.Upload, func(r rc_recipe.Recipe) {
		ru := r.(*file.Upload)
		ru.EstimateOnly = true
		ru.LocalPath = z.LocalPath
		ru.DropboxPath = z.DropboxPath
		ru.Context = z.Peer.Context()
	})
}

func (z *Up) Test(c app_control.Control) error {
	return qt_endtoend.ScenarioTest()
}
