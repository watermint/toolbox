package daily

import (
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_activity"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_time"
)

type Event struct {
	Peer      dbx_conn.ConnBusinessAudit
	StartDate string
	EndDate   mo_string.OptionalString
	Category  mo_string.OptionalString
	Event     rp_model.RowReport
	Progress  app_msg.Message
}

func (z *Event) Preset() {
	z.Event.SetModel(&mo_activity.Compatible{})
}

func (z *Event) Exec(c app_control.Control) error {
	ui := c.UI()

	dr, err := ut_time.Daily(z.StartDate, z.EndDate.Value())
	if err != nil {
		return err
	}

	for _, d := range dr {
		st, _ := ut_time.ParseTimestamp(d.Start)
		stDate := st.Format("2006-01-02")
		ui.Progress(z.Progress.With("Start", d.Start).With("End", d.End))

		rep, err := z.Event.OpenNew(rp_model.Suffix(stDate))
		if err != nil {
			return err
		}

		handler := func(event *mo_activity.Event) error {
			rep.Row(event.Compatible())
			return nil
		}

		err = sv_activity.New(z.Peer.Context()).List(handler,
			sv_activity.StartTime(d.Start),
			sv_activity.EndTime(d.End),
			sv_activity.Category(z.Category.Value()),
		)
		rep.Close()
		if err != nil {
			return err
		}
	}

	return err
}

func (z *Event) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Event{}, func(r rc_recipe.Recipe) {
		m := r.(*Event)
		m.StartDate = "2020-03-10"
	})
}
