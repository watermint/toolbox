package activity

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_activity"
	"github.com/watermint/toolbox/domain/service/sv_activity"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"time"
)

type EventVO struct {
	Peer      rc_conn.ConnBusinessAudit
	StartTime string
	EndTime   string
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

func (z *Event) Requirement() rc_vo.ValueObject {
	return &EventVO{}
}

func (z *Event) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*EventVO)
	l := k.Log()

	if vo.StartTime != "" {
		if t, ok := ut_time.ParseTimestamp(vo.StartTime); ok {
			l.Debug("Rebase StartTime", zap.String("startTime", vo.StartTime))
			vo.StartTime = api_util.RebaseAsString(t)
			l.Debug("Rebased StartTime", zap.String("startTime", vo.StartTime))
		} else {
			return errors.New("invalid date/time format for -start-date")
		}
	}
	if vo.EndTime != "" {
		if t, ok := ut_time.ParseTimestamp(vo.EndTime); ok {
			l.Debug("Rebase EndTime", zap.String("endTime", vo.StartTime))
			vo.StartTime = api_util.RebaseAsString(t)
			l.Debug("Rebased EndTime", zap.String("endTime", vo.StartTime))
		} else {
			return errors.New("invalid date/time format for -end-date")
		}
	}

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportEvent)
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
	if !qt_recipe.ApplyTestPeers(c, lvo) {
		return qt_recipe.NotEnoughResource()
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "event", func(cols map[string]string) error {
		if _, ok := cols["timestamp"]; !ok {
			return errors.New("`timestamp` is not found")
		}
		return nil
	})
}
