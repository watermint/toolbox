package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/file/es_size"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Size struct {
	Peer     dbx_conn.ConnScopedIndividual
	Size     rp_model.RowReport
	Path     mo_path.DropboxPath
	Depth    mo_int.RangeInt
	Folder   kv_storage.Storage
	Sum      kv_storage.Storage
	BasePath mo_string.SelectString
}

func (z *Size) Preset() {
	z.Size.SetModel(
		&es_size.FolderSize{},
	)
	z.Depth.SetRange(1, 300, 2)
	z.Peer.SetScopes(dbx_auth.ScopeFilesContentRead)
	z.Path = mo_path.NewDropboxPath("/")
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Size) Exec(c app_control.Control) error {
	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	fs := dbx_fs.NewFileSystem(client)

	factory := c.NewKvsFactory()
	defer func() {
		factory.Close()
	}()

	if err := z.Size.Open(); err != nil {
		return err
	}

	return es_size.ScanSingleFileSystem(
		c.Log(),
		c.Sequence(),
		z.Folder,
		z.Sum,
		fs,
		dbx_fs.NewPath("", z.Path),
		z.Depth.Value(),
		func(s es_size.FolderSize) {
			z.Size.Row(&s)
		},
	)
}

func (z *Size) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Size{}, func(r rc_recipe.Recipe) {
		m := r.(*Size)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath()
	})
}
