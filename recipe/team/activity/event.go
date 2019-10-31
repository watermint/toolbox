package activity

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_activity"
	"github.com/watermint/toolbox/domain/service/sv_activity"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"time"
)

type EventVO struct {
	Peer      app_conn.ConnBusinessAudit
	StartTime string
	EndTime   string
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

	rep, err := k.Report("event", &mo_activity.Event{})
	if err != nil {
		return err
	}
	defer rep.Close()

	handler := func(event *mo_activity.Event) error {
		rep.Row(event)
		return nil
	}

	return sv_activity.New(ctx).List(handler,
		sv_activity.StartTime(vo.StartTime),
		sv_activity.EndTime(vo.EndTime),
		sv_activity.Category(vo.Category),
	)
}

func (z *Event) Test(c app_control.Control) error {
	lvo := &EventVO{
		StartTime: api_util.RebaseAsString(time.Now().Add(-10 * time.Minute)),
		Category:  "logins",
	}
	if !app_test.ApplyTestPeers(c, lvo) {
		return qt_test.NotEnoughResource()
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "event", func(cols map[string]string) error {
		if _, ok := cols["timestamp"]; !ok {
			return errors.New("`timestamp` is not found")
		}
		return nil
	})
}
