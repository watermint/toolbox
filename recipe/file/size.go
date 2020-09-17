package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/file/es_size"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Size struct {
	Peer   dbx_conn.ConnUserFile
	Size   rp_model.RowReport
	Errors rp_model.TransactionReport
	Path   mo_path.DropboxPath
	Depth  mo_int.RangeInt
}

func (z *Size) Preset() {
	z.Size.SetModel(
		&es_size.FolderSize{},
	)
	z.Depth.SetRange(1, 300, 2)
	z.Path = mo_path.NewDropboxPath("/")
}

func (z *Size) Exec(c app_control.Control) error {
	fs := filesystem.NewFileSystem(z.Peer.Context())

	traverse, err := es_size.New(c.Log(), kv_storage_impl.NewFactory(c), c.Sequence(), fs, z.Depth.Value())
	if err != nil {
		return err
	}

	if err := z.Errors.Open(rp_model.NoConsoleOutput()); err != nil {
		return err
	}
	if err := z.Size.Open(); err != nil {
		return err
	}

	return traverse.Scan(filesystem.NewPath("", z.Path), func(s es_size.FolderSize) {
		z.Size.Row(&s)
	})
}

func (z *Size) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Size{}, func(r rc_recipe.Recipe) {
		m := r.(*Size)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath()
	})
}
