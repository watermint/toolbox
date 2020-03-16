package file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file_content"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type Download struct {
	Peer         rc_conn.ConnUserFile
	DropboxPath  mo_path.DropboxPath
	LocalPath    mo_path.FileSystemPath
	OperationLog rp_model.RowReport
}

func (z *Download) Preset() {
	z.OperationLog.SetModel(&mo_file.ConcreteEntry{})
}

func (z *Download) Exec(c app_control.Control) error {
	l := c.Log()
	ctx := z.Peer.Context()

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	entry, f, err := sv_file_content.NewDownload(ctx).Download(z.DropboxPath)
	if err != nil {
		return err
	}
	if err := os.Rename(f.Path(), filepath.Join(z.LocalPath.Path(), entry.Name())); err != nil {
		l.Debug("Unable to move file to specified path",
			zap.Error(err),
			zap.String("downloaded", f.Path()),
			zap.String("destination", z.LocalPath.Path()),
		)
		return err
	}

	z.OperationLog.Row(entry.Concrete())
	return nil
}

func (z *Download) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Download{}, func(r rc_recipe.Recipe) {
		m := r.(*Download)
		m.LocalPath = qt_recipe.NewTestFileSystemFolderPath(c, "download")
		m.DropboxPath = qt_recipe.NewTestDropboxFolderPath("file-download")
	})
}
