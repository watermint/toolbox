package sync

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_filter"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	ig_file2 "github.com/watermint/toolbox/ingredient/ig_dropbox/ig_file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Online struct {
	rc_recipe.RemarkIrreversible
	Peer         dbx_conn.ConnScopedIndividual
	Src          mo_path.DropboxPath
	Dst          mo_path.DropboxPath
	Upload       *ig_file2.Upload
	SkipExisting bool
	Delete       bool
	Name         mo_filter.Filter
	Online       *ig_file2.Online
	BasePath     mo_string.SelectString
}

func (z *Online) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
	)
	z.Name.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNameSuffixFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_file_filter.NewIgnoreFileFilter(),
	)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Online) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.Online, func(r rc_recipe.Recipe) {
		ru := r.(*ig_file2.Online)
		ru.SrcPath = z.Src
		ru.DstPath = z.Dst
		ru.Overwrite = !z.SkipExisting
		ru.Name = z.Name
		ru.Context = z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	})
}

func (z *Online) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Online{}, func(r rc_recipe.Recipe) {
		m := r.(*Online)
		m.Src = qtr_endtoend.NewTestDropboxFolderPath("src")
		m.Dst = qtr_endtoend.NewTestDropboxFolderPath("dst")
	})
	if err, _ = qt_errors.ErrorsForTest(c.Log(), err); err != nil {
		return err
	}

	return qt_errors.ErrorScenarioTest
}
