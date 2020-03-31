package activity

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_activity"
	"github.com/watermint/toolbox/infra/api/dbx_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"time"
)

type Event struct {
	Peer      rc_conn.ConnBusinessAudit
	StartTime mo_time.TimeOptional
	EndTime   mo_time.TimeOptional
	Category  string
	Event     rp_model.RowReport
}

func (z *Event) Preset() {
	z.Event.SetModel(&mo_activity.Compatible{})
}

func (z *Event) Exec(c app_control.Control) error {
	if err := z.Event.Open(); err != nil {
		return err
	}

	handler := func(event *mo_activity.Event) error {
		z.Event.Row(event.Compatible())
		return nil
	}

	return sv_activity.New(z.Peer.Context()).List(handler,
		sv_activity.StartTime(z.StartTime.Iso8601()),
		sv_activity.EndTime(z.EndTime.Iso8601()),
		sv_activity.Category(z.Category),
	)
}

func (z *Event) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &Event{}, func(r rc_recipe.Recipe) {
		rc := r.(*Event)
		if t, ok := rc.StartTime.(*mo_time.TimeImpl); ok {
			t.UpdateTime(time.Now().Add(-10 * time.Minute).Format(dbx_util.DateTimeFormat))
		}
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
