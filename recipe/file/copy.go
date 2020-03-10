package file

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_relocation"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type Copy struct {
	Peer rc_conn.ConnUserFile
	Src  mo_path.DropboxPath
	Dst  mo_path.DropboxPath
}

func (z *Copy) Preset() {
}

func (z *Copy) Exec(c app_control.Control) error {
	uc := uc_file_relocation.New(z.Peer.Context())
	return uc.Copy(z.Src, z.Dst)
}

func (z *Copy) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Copy{}, func(r rc_recipe.Recipe) {
		m := r.(*Copy)
		m.Src = qt_recipe.NewTestDropboxFolderPath("src")
		m.Dst = qt_recipe.NewTestDropboxFolderPath("dst")
	})
	if err, _ = qt_recipe.RecipeError(c.Log(), err); err != nil {
		return err
	}
	return qt_errors.ErrorScenarioTest
}
