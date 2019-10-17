package batch

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file_url"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type UrlVO struct {
	Peer app_conn.ConnUserFile
	Data app_file.Data
}

type UrlRow struct {
	Url  string
	Path string
}

type UrlWorker struct {
	url  string
	path string
	ctx  api_context.Context
	ctl  app_control.Control
	rep  app_report.Report
}

func (z *UrlWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Info("recipe.file.import.batch.url.progress", app_msg.P{
		"Url":  z.url,
		"Path": z.path,
	})

	entry, err := sv_file_url.New(z.ctx).Save(mo_path.NewPath(z.path), z.url)
	if err != nil {
		return err
	}
	z.rep.Row(entry.Concrete())

	return nil
}

type Url struct {
}

func (z *Url) Requirement() app_vo.ValueObject {
	return &UrlVO{}
}

func (z *Url) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*UrlVO)

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	rep, err := k.Report("import_url", &mo_file.ConcreteEntry{})
	if err != nil {
		return err
	}
	defer rep.Close()

	err = vo.Data.Model(k.Control(), &UrlRow{})
	if err != nil {
		return err
	}

	q := k.NewQueue()
	err = vo.Data.EachRow(func(m interface{}, rowIndex int) error {
		r := m.(*UrlRow)
		q.Enqueue(&UrlWorker{
			url:  r.Url,
			path: r.Path,
			ctx:  ctx,
			ctl:  k.Control(),
			rep:  rep,
		})
		return nil
	})
	q.Wait()
	return nil
}

func (z *Url) Test(c app_control.Control) error {
	return nil
}
