package event

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/calendar/model/mo_event"
	"github.com/watermint/toolbox/domain/google/calendar/service/sv_event"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/time/ut_compare"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"time"
)

type List struct {
	Peer        goog_conn.ConnGoogleCalendar
	CalendarId  mo_string.OptionalString
	Events      rp_model.RowReport
	Query       mo_string.OptionalString
	Start       mo_time.TimeOptional
	End         mo_time.TimeOptional
	DoNotFilter bool
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeCalendarEventsReadOnly,
	)
	z.Events.SetModel(
		&mo_event.Event{},
		rp_model.HiddenColumns(
			"id",
			"description",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Events.Open(); err != nil {
		return err
	}
	var calendarId string
	if z.CalendarId.IsExists() {
		calendarId = z.CalendarId.Value()
	} else {
		calendarId = "primary"
	}
	opts := make([]func(o sv_event.ListOpts) sv_event.ListOpts, 0)
	if z.Query.IsExists() {
		opts = append(opts, sv_event.Query(z.Query.Value()))
	}
	if !z.Start.IsZero() {
		opts = append(opts, sv_event.TimeMin(z.Start.Iso8601()))
	} else {
		// start from -30days
		opts = append(opts, sv_event.TimeMin(dbx_util.ToApiTimeString(time.Now().Add(-30*24*time.Hour))))
	}
	if !z.End.IsZero() {
		opts = append(opts, sv_event.TimeMax(z.End.Iso8601()))
	}

	return sv_event.New(z.Peer.Client()).ListEach(func(event *mo_event.Event) {
		if z.DoNotFilter {
			z.Events.Row(event)
			return
		}

		switch {
		case event.StartDate != "":
			startDate, err := dbx_util.Parse(event.StartDate)
			// do not filter in case of error
			if err != nil {
				z.Events.Row(event)
				return
			}
			if ut_compare.IsBetweenOptional(startDate, z.Start, z.End) {
				z.Events.Row(event)
			}

		case event.StartDateTime != "":
			startDateTime, err := dbx_util.Parse(event.StartDate)
			// do not filter in case of error
			if err != nil {
				z.Events.Row(event)
				return
			}
			if ut_compare.IsBetweenOptional(startDateTime, z.Start, z.End) {
				z.Events.Row(event)
			}
		}
	}, calendarId, lang.ApplyOpts(sv_event.ListOpts{}, opts))
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
