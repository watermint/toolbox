package batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_lock"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type PathLock struct {
	Path string `json:"path"`
}

type Acquire struct {
	Peer         dbx_conn.ConnScopedIndividual
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
	BatchSize    int
}

func (z *Acquire) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
	)
	z.BatchSize = 100
	z.File.SetModel(&PathLock{})
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

func (z *Acquire) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}
	sfl := sv_file_lock.New(z.Peer.Context())
	var lastErr error

	lockBucket := func(bucket []mo_path.DropboxPath) {
		if len(bucket) < 1 {
			l.Debug("No more bucket")
			return
		}

		locked, err := sfl.LockBatch(bucket)
		if err != nil {
			l.Debug("Something happened during lock batch", esl.Error(err))
			lastErr = err
		}
		for path, info := range locked {
			if info.Error != nil {
				z.OperationLog.Failure(info.Error, PathLock{
					Path: path,
				})
			} else {
				z.OperationLog.Success(PathLock{
					Path: path,
				}, info.Entry.LockInfo())
			}
		}
	}

	bucket := make([]mo_path.DropboxPath, 0)
	loadErr := z.File.EachRow(func(m interface{}, rowIndex int) error {
		row := m.(*PathLock)
		bucket = append(bucket, mo_path.NewDropboxPath(row.Path))
		if z.BatchSize <= len(bucket) {
			l.Debug("Bucket exceeds batch size, lock those", esl.Int("bucketLen", len(bucket)))
			lockBucket(bucket)
			bucket = make([]mo_path.DropboxPath, 0)
		}
		return nil
	})
	// lock remaining
	lockBucket(bucket)

	if loadErr != nil {
		return loadErr
	}

	return lastErr
}

func (z *Acquire) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("lock", "/Test/a.txt\n/Test/b.txt")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecMock(c, &Acquire{}, func(r rc_recipe.Recipe) {
		m := r.(*Acquire)
		m.File.SetFilePath(f)
	})
}
