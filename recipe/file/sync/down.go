package sync

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_filter"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/ig_file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Down struct {
	Peer         dbx_conn.ConnScopedIndividual
	Delete       bool
	SkipExisting bool
	LocalPath    mo_path2.FileSystemPath
	DropboxPath  mo_path.DropboxPath
	Name         mo_filter.Filter
	Download     *ig_file.Download
}

func (z *Down) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
	)
	z.Name.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNameSuffixFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_file_filter.NewIgnoreFileFilter(),
	)
}

func (z *Down) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.Download, func(r rc_recipe.Recipe) {
		ru := r.(*ig_file.Download)
		ru.LocalPath = z.LocalPath
		ru.DropboxPath = z.DropboxPath
		ru.Overwrite = !z.SkipExisting
		ru.Name = z.Name
		ru.Context = z.Peer.Client()
	})
}

func (z *Down) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Down{}, func(r rc_recipe.Recipe) {
		m := r.(*Down)
		m.LocalPath = qtr_endtoend.NewTestExistingFileSystemFolderPath(c, "down")
		m.DropboxPath = qtr_endtoend.NewTestDropboxFolderPath("down")
	})
	if err, _ = qt_errors.ErrorsForTest(c.Log(), err); err != nil {
		return err
	}

	return qt_errors.ErrorScenarioTest
}
