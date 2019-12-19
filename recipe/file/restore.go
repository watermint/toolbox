package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_file_restore"
	"github.com/watermint/toolbox/domain/service/sv_file_revision"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type RestoreVO struct {
	Peer rc_conn.ConnUserFile
	Path string
}

type RestoreWorker struct {
	k    rc_kitchen.Kitchen
	ctx  api_context.Context
	rep  rp_model.Report
	path mo_path.DropboxPath
}

type RestoreTarget struct {
	Path string `json:"path"`
}

func (z *RestoreWorker) Exec() error {
	l := z.k.Log().With(zap.String("path", z.path.Path()))
	ui := z.k.UI()
	ui.Info("recipe.file.restore.progress.restore_file", app_msg.P{"Path": z.path.Path()})
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
		z.rep.Skip(app_msg.M("recipe.file.restore.skip.is_not_deleted"), target)
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

const (
	reportRestore = "restore"
)

type Restore struct {
}

func (z *Restore) Console() {
}

func (z *Restore) Requirement() rc_vo.ValueObject {
	return &RestoreVO{}
}

func (z *Restore) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*RestoreVO)
	ui := k.UI()
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportRestore)
	if err != nil {
		return err
	}
	defer rep.Close()
	q := k.NewQueue()

	count := 0
	handler := func(entry mo_file.Entry) {
		if f, e := entry.Deleted(); e {
			count++
			q.Enqueue(&RestoreWorker{
				k:    k,
				ctx:  ctx,
				path: f.Path(),
				rep:  rep,
			})
		}
	}

	lastErr := sv_file.NewFiles(ctx).ListChunked(
		mo_path.NewDropboxPath(vo.Path),
		handler,
		sv_file.IncludeDeleted(),
		sv_file.Recursive(),
	)
	q.Wait()

	ui.Info("recipe.file.restore.progress.finish", app_msg.P{
		"Count": count,
	})

	return lastErr
}

func (z *Restore) Test(c app_control.Control) error {
	return qt_recipe.ImplementMe()
}

func (z *Restore) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportRestore,
			rp_model.TransactionHeader(
				&RestoreTarget{},
				&mo_file.ConcreteEntry{},
			),
		),
	}
}
