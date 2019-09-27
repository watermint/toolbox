package team

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
	"time"
)

type ActivityVO struct {
	PeerName  app_conn.ConnBusinessAudit
	StartTime string
	EndTime   string
	Category  string
}

type Activity struct {
}

func (z *Activity) Requirement() app_vo.ValueObject {
	return &ActivityVO{}
}

func (z *Activity) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*ActivityVO)

	ctx, err := vo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	rep, err := k.Report("activity", &mo_activity.Event{})
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

func (z *Activity) Test(c app_control.Control) error {
	lvo := &ActivityVO{
		StartTime: api_util.RebaseAsString(time.Now().Add(-10 * time.Minute)),
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
