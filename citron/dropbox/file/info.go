package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Info struct {
	Peer     dbx_conn.ConnScopedIndividual
	Path     mo_path.DropboxPath
	Metadata rp_model.RowReport
}

func (z *Info) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
	)
	z.Metadata.SetModel(
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"id",
			"path_lower",
		),
	)
}

func (z *Info) Exec(c app_control.Control) error {
	if err := z.Metadata.Open(); err != nil {
		return err
	}

	entry, err := sv_file.NewFiles(z.Peer.Client()).Resolve(z.Path)
	if err != nil {
		return err
	}

	z.Metadata.Row(entry.Concrete())
	return nil
}

func (z *Info) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Info{}, func(r rc_recipe.Recipe) {
		m := r.(*Info)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("files-info")
	})
}
