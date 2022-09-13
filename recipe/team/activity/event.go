package activity

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_activity"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"time"
)

type Event struct {
	Peer      dbx_conn.ConnScopedTeam
	StartTime mo_time.TimeOptional
	EndTime   mo_time.TimeOptional
	Category  mo_string.OptionalString
	Event     rp_model.RowReport
}

func (z *Event) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeEventsRead,
	)
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

	opts := make([]sv_activity.ListOpt, 0)
	opts = append(opts, sv_activity.StartTime(z.StartTime.Iso8601()))
	opts = append(opts, sv_activity.EndTime(z.EndTime.Iso8601()))
	if z.Category.IsExists() {
		opts = append(opts, sv_activity.Category(z.Category.Value()))
	}

	return sv_activity.New(z.Peer.Client()).List(handler, opts...)
}

func (z *Event) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &Event{}, func(r rc_recipe.Recipe) {
		rc := r.(*Event)
		if t, ok := rc.StartTime.(*mo_time.TimeImpl); ok {
			t.UpdateTime(time.Now().Add(-10 * time.Minute).Format(dbx_util.DateTimeFormat))
		}
		rc.Category = mo_string.NewOptional("logins")
	})
	if err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "event", func(cols map[string]string) error {
		if _, ok := cols["timestamp"]; !ok {
			return errors.New("`timestamp` is not found")
		}
		return nil
	})
}
