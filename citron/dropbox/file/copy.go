package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_relocation"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Copy struct {
	Peer     dbx_conn.ConnScopedIndividual
	Src      mo_path.DropboxPath
	Dst      mo_path.DropboxPath
	BasePath mo_string.SelectString
}

func (z *Copy) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
	)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Copy) Exec(c app_control.Control) error {
	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))

	uc := uc_file_relocation.New(client)
	return uc.Copy(z.Src, z.Dst)
}

func (z *Copy) Test(c app_control.Control) error {
	// replay test
	{
		err := rc_exec.ExecReplay(c, &Copy{}, "recipe-file-copy.json.gz", func(r rc_recipe.Recipe) {
			m := r.(*Copy)
			m.Src = qtr_endtoend.NewTestDropboxFolderPath("src")
			m.Dst = qtr_endtoend.NewTestDropboxFolderPath("dst")
		})
		if err, _ = qt_errors.ErrorsForTest(c.Log(), err); err != nil {
			return err
		}
	}

	err := rc_exec.ExecMock(c, &Copy{}, func(r rc_recipe.Recipe) {
		m := r.(*Copy)
		m.Src = qtr_endtoend.NewTestDropboxFolderPath("src")
		m.Dst = qtr_endtoend.NewTestDropboxFolderPath("dst")
	})
	if err, _ = qt_errors.ErrorsForTest(c.Log(), err); err != nil {
		return err
	}
	return qt_errors.ErrorScenarioTest
}
