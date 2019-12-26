package batch

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file_url"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type UrlRow struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

type UrlWorker struct {
	row *UrlRow
	ctx api_context.Context
	ctl app_control.Control
	rep rp_model.TransactionReport
}

func (z *UrlWorker) Exec() error {
	ui := z.ctl.UI()

	path := sv_file_url.PathWithName(mo_path.NewDropboxPath(z.row.Path), z.row.Url)
	ui.Info("recipe.file.import.batch.url.progress", app_msg.P{
		"Url":  z.row.Url,
		"Path": path.Path(),
	})

	entry, err := sv_file_url.New(z.ctx).Save(path, z.row.Url)
	if err != nil {
		z.rep.Failure(err, z.row)
		return err
	}
	z.rep.Success(z.row, entry.Concrete())

	return nil
}

const (
	reportUrl = "import_url"
)

type Url struct {
	Peer            rc_conn.ConnUserFile
	File            fd_file.RowFeed
	Path            string
	OperationLog    rp_model.TransactionReport
	SkipPathMissing app_msg.Message
}

func (z *Url) Preset() {
	z.OperationLog.SetModel(&UrlRow{}, &mo_file.ConcreteEntry{})
	z.File.SetModel(&UrlRow{})
}

func (z *Url) Exec(k rc_kitchen.Kitchen) error {
	ui := k.UI()
	ctx := z.Peer.Context()

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	q := k.NewQueue()
	err := z.File.EachRow(func(m interface{}, rowIndex int) error {
		r := m.(*UrlRow)
		var path string
		switch {
		case r.Path != "":
			path = r.Path
		case z.Path != "":
			path = z.Path
		default:
			z.OperationLog.Skip(z.SkipPathMissing, r)
			ui.Error("recipe.file.import.batch.url.err.path_missing")
			return errors.New("no path to save")
		}

		q.Enqueue(&UrlWorker{
			row: &UrlRow{
				Url:  r.Url,
				Path: path,
			},
			ctx: ctx,
			ctl: k.Control(),
			rep: z.OperationLog,
		})
		return nil
	})
	q.Wait()
	return err
}

func (z *Url) Test(c app_control.Control) error {
	return qt_recipe.ImplementMe()
}
