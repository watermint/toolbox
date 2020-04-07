package export

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type Doc struct {
	Peer         dbx_conn.ConnUserFile
	LocalPath    mo_path.FileSystemPath
	DropboxPath  mo_path.DropboxPath
	OperationLog rp_model.RowReport
}

func (z *Doc) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	export, path, err := sv_file_content.NewExport(z.Peer.Context()).Export(z.DropboxPath)
	if err != nil {
		return err
	}
	dest := filepath.Join(z.LocalPath.Path(), export.ExportName)
	if err := os.Rename(path.Path(), dest); err != nil {
		l.Debug("Unable to move file to specified path",
			zap.Error(err),
			zap.String("downloaded", path.Path()),
			zap.String("destination", dest),
		)
		return err
	}

	z.OperationLog.Row(export)

	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Doc{}, func(r rc_recipe.Recipe) {
		m := r.(*Doc)
		m.LocalPath = qt_recipe.NewTestFileSystemFolderPath(c, "export-doc")
		m.DropboxPath = qt_recipe.NewTestDropboxFolderPath("file-export-doc")
	})
}

func (z *Doc) Preset() {
	z.OperationLog.SetModel(
		&mo_file.Export{},
		rp_model.HiddenColumns(
			"path_lower",
			"id",
			"revision",
			"content_hash",
			"export_hash",
		),
	)
}
