package activity

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_activity"
	"github.com/watermint/toolbox/domain/service/sv_activity"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"time"
)

type Event struct {
	Peer      rc_conn.ConnBusinessAudit
	StartTime string
	EndTime   string
	Category  string
	Event     rp_model.RowReport
}

func (z *Event) Preset() {
	z.Event.SetModel(&mo_activity.Compatible{})
}

func (z *Event) Exec(c app_control.Control) error {
	l := c.Log()

	if z.StartTime != "" {
		if t, ok := ut_time.ParseTimestamp(z.StartTime); ok {
			l.Debug("Rebase StartTime", zap.String("startTime", z.StartTime))
			z.StartTime = api_util.RebaseAsString(t)
			l.Debug("Rebased StartTime", zap.String("startTime", z.StartTime))
		} else {
			return errors.New("invalid date/time format for -start-date")
		}
	}
	if z.EndTime != "" {
		if t, ok := ut_time.ParseTimestamp(z.EndTime); ok {
			l.Debug("Rebase EndTime", zap.String("endTime", z.StartTime))
			z.StartTime = api_util.RebaseAsString(t)
			l.Debug("Rebased EndTime", zap.String("endTime", z.StartTime))
		} else {
			return errors.New("invalid date/time format for -end-date")
		}
	}

	if err := z.Event.Open(); err != nil {
		return err
	}

	handler := func(event *mo_activity.Event) error {
		z.Event.Row(event.Compatible())
		return nil
	}

	return sv_activity.New(z.Peer.Context()).List(handler,
		sv_activity.StartTime(z.StartTime),
		sv_activity.EndTime(z.EndTime),
		sv_activity.Category(z.Category),
	)
}

func (z *Event) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &Event{}, func(r rc_recipe.Recipe) {
		rc := r.(*Event)
		rc.StartTime = api_util.RebaseAsString(time.Now().Add(-10 * time.Minute))
		rc.Category = "logins"
	})
	if err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "event", func(cols map[string]string) error {
		if _, ok := cols["timestamp"]; !ok {
			return errors.New("`timestamp` is not found")
		}
		return nil
	})
}
