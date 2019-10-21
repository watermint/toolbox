package daily

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_activity"
	"github.com/watermint/toolbox/domain/service/sv_activity"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"time"
)

type EventVO struct {
	Peer      app_conn.ConnBusinessAudit
	StartDate string
	EndDate   string
	Category  string
}

type Event struct {
}

func (z *Event) Requirement() app_vo.ValueObject {
	return &EventVO{}
}

func (z *Event) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*EventVO)

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

		rep, err := k.Report("event"+stDate, &mo_activity.Event{})
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
	lvo := &EventVO{
		StartDate: api_util.RebaseAsString(time.Now().Add(-10 * time.Minute)),
		Category:  "logins",
	}
	if !app_test.ApplyTestPeers(c, lvo) {
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "activity", func(cols map[string]string) error {
		if _, ok := cols["timestamp"]; !ok {
			return errors.New("`timestamp` is not found")
		}
		return nil
	})
}
