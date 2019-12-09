package daily

import (
	"github.com/watermint/toolbox/domain/model/mo_activity"
	"github.com/watermint/toolbox/domain/service/sv_activity"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type EventVO struct {
	Peer      app_conn.ConnBusinessAudit
	StartDate string
	EndDate   string
	Category  string
}

const (
	reportEvent = "event"
)

type Event struct {
}

func (z *Event) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportEvent, &mo_activity.Event{}),
	}
}

func (z *Event) Requirement() app_vo.ValueObject {
	return &EventVO{}
}

func (z *Event) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*EventVO)
	ui := k.UI()

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	dr, err := ut_time.Daily(vo.StartDate, vo.EndDate)
	if err != nil {
		return err
	}

	for _, d := range dr {
		st, _ := ut_time.ParseTimestamp(d.Start)
		stDate := st.Format("2006-01-02")

		ui.Info("recipe.team.activity.daily.event.progress", app_msg.P{
			"Start": d.Start,
			"End":   d.End,
		})

		rep, err := rp_spec_impl.New(z, k.Control()).Open(reportEvent, rp_model.Suffix(stDate))
		if err != nil {
			return err
		}

		handler := func(event *mo_activity.Event) error {
			rep.Row(event)
			return nil
		}

		err = sv_activity.New(ctx).List(handler,
			sv_activity.StartTime(d.Start),
			sv_activity.EndTime(d.End),
			sv_activity.Category(vo.Category),
		)
		rep.Close()
	}

	return err
}

func (z *Event) Test(c app_control.Control) error {
	return qt_recipe.ImplementMe()
}
