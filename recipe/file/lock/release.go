package lock

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_lock"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Release struct {
	Peer         dbx_conn.ConnScopedIndividual
	Path         mo_path.DropboxPath
	OperationLog rp_model.TransactionReport
}

func (z *Release) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
	)
	z.OperationLog.SetModel(
		&PathLock{},
		&mo_file.LockInfo{},
		rp_model.HiddenColumns(
			"result.id",
			"result.name",
			"result.path_lower",
			"result.path_display",
			"result.revision",
			"result.content_hash",
			"result.shared_folder_id",
			"result.parent_shared_folder_id",
			"result.lock_holder_account_id",
		),
	)
}

func (z *Release) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	entry, err := sv_file_lock.New(z.Peer.Context()).Unlock(z.Path)
	if err != nil {
		z.OperationLog.Failure(err, &PathLock{Path: z.Path.Path()})
		return err
	}
	z.OperationLog.Success(&PathLock{Path: z.Path.Path()}, entry.LockInfo())
	return nil
}

func (z *Release) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Release{}, func(r rc_recipe.Recipe) {
		m := r.(*Release)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("test.txt")
	})
}
