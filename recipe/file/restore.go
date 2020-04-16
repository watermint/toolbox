package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_restore"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_revision"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type MsgRestore struct {
	ProgressRestore  app_msg.Message
	ProgressFinish   app_msg.Message
	SkipIsNotDeleted app_msg.Message
}

var (
	MRestore = app_msg.Apply(&MsgRestore{}).(*MsgRestore)
)

type RestoreWorker struct {
	ctl  app_control.Control
	ctx  dbx_context.Context
	rep  rp_model.TransactionReport
	path mo_path.DropboxPath
}

type RestoreTarget struct {
	Path string `json:"path"`
}

func (z *RestoreWorker) Exec() error {
	l := z.ctl.Log().With(zap.String("path", z.path.Path()))
	ui := z.ctl.UI()
	ui.Progress(MRestore.ProgressRestore.With("Path", z.path.Path()))
	target := &RestoreTarget{
		Path: z.path.Path(),
	}

	revs, err := sv_file_revision.New(z.ctx).List(z.path)
	if err != nil {
		l.Debug("Unable to retrieve revisions", zap.Error(err))
		z.rep.Failure(err, target)
		return err
	}
	if !revs.IsDeleted {
		l.Debug("The file is not deleted")
		z.rep.Skip(MRestore.SkipIsNotDeleted, target)
		return nil
	}
	if len(revs.Entries) < 1 {
		l.Debug("No revision found")
		err = errors.New("no revisions found for the file")
		z.rep.Failure(err, target)
		return err
	}
	targetRev := revs.Entries[0].Revision
	l.Debug("Restoring to most recent state", zap.String("targetRev", targetRev))

	e, err := sv_file_restore.New(z.ctx).Restore(z.path, targetRev)
	if err != nil {
		z.rep.Failure(err, target)
		return err
	}
	z.rep.Success(target, e.Concrete())
	return nil
}

type Restore struct {
	Peer         dbx_conn.ConnUserFile
	Path         mo_path.DropboxPath
	OperationLog rp_model.TransactionReport
}

func (z *Restore) Preset() {
	z.OperationLog.SetModel(
		&RestoreTarget{},
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"result.id",
			"result.path_lower",
			"result.revision",
			"result.content_hash",
			"result.shared_folder_id",
			"result.parent_shared_folder_id",
		),
	)
}

func (z *Restore) Exec(c app_control.Control) error {
	ui := c.UI()
	ctx := z.Peer.Context()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	q := c.NewQueue()

	count := 0
	handler := func(entry mo_file.Entry) {
		if f, e := entry.Deleted(); e {
			count++
			q.Enqueue(&RestoreWorker{
				ctl:  c,
				ctx:  ctx,
				path: f.Path(),
				rep:  z.OperationLog,
			})
		}
	}

	lastErr := sv_file.NewFiles(ctx).ListChunked(
		z.Path,
		handler,
		sv_file.IncludeDeleted(),
		sv_file.Recursive(),
	)
	q.Wait()

	ui.Info(MRestore.ProgressFinish.With("Count", count))

	return lastErr
}

func (z *Restore) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Restore{}, func(r rc_recipe.Recipe) {
		m := r.(*Restore)
		m.Path = qt_recipe.NewTestDropboxFolderPath("file-restore")
	})
}
