package daily

import (
	"github.com/watermint/toolbox/domain/model/mo_activity"
	"github.com/watermint/toolbox/domain/service/sv_activity"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Event struct {
	Peer      rc_conn.ConnBusinessAudit
	StartDate string
	EndDate   string
	Category  string
	Event     rp_model.RowReport
}

func (z *Event) Preset() {
	z.Event.SetModel(&mo_activity.Compatible{})
}

func (z *Event) Exec(c app_control.Control) error {
	ui := c.UI()

	dr, err := ut_time.Daily(z.StartDate, z.EndDate)
	if err != nil {
		return err
	}

	for _, d := range dr {
		st, _ := ut_time.ParseTimestamp(d.Start)
		stDate := st.Format("2006-01-02")

		ui.InfoK("recipe.team.activity.daily.event.progress", app_msg.P{
			"Start": d.Start,
			"End":   d.End,
		})

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
			sv_activity.Category(z.Category),
		)
		rep.Close()
		if err != nil {
			return err
		}
	}

	return err
}

func (z *Event) Test(c app_control.Control) error {
	return qt_errors.ErrorImplementMe
}
