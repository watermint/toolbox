package batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Release struct {
	Peer         dbx_conn.ConnScopedIndividual
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
}

func (z *Release) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
	)
	z.File.SetModel(&PathLock{})
	z.OperationLog.SetModel(
		&PathLock{},
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"result.id",
			"result.name",
			"result.path_lower",
			"result.path_display",
			"result.revision",
			"result.content_hash",
			"result.shared_folder_id",
			"result.parent_shared_folder_id",
		),
	)
}

func (z *Release) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	var lastErr error
	_ = z.File.EachRow(func(m interface{}, rowIndex int) error {
		row := m.(*PathLock)
		entry, err := sv_file.NewFiles(z.Peer.Context()).Unlock(mo_path.NewDropboxPath(row.Path))
		if err != nil {
			z.OperationLog.Failure(err, &PathLock{Path: row.Path})
			lastErr = err
			return nil
		}
		z.OperationLog.Success(&PathLock{Path: row.Path}, entry.Concrete())
		return nil
	})
	return lastErr
}

func (z *Release) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("lock", "/Test/a.txt\n/Test/b.txt")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecMock(c, &Release{}, func(r rc_recipe.Recipe) {
		m := r.(*Release)
		m.File.SetFilePath(f)
	})
}
