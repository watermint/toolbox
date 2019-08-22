package sv_activity

import (
	"github.com/watermint/toolbox/domain/model/mo_activity"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
)

type Activity interface {
	All(handler func(event *mo_activity.Event) error) (err error)
}

func New(ctx api_context.Context) Activity {
	return &activityImpl{
		ctx: ctx,
	}
}

type activityImpl struct {
	ctx api_context.Context
}

func (z *activityImpl) All(handler func(event *mo_activity.Event) error) (err error) {
	return z.ctx.List("team_log/get_events").
		Continue("team_log/get_events/continue").
		UseHasMore(true).
		ResultTag("events").
		OnEntry(func(entry api_list.ListEntry) error {
			e := &mo_activity.Event{}
			if err = entry.Model(e); err != nil {
				return err
			}
			return handler(e)
		}).Call()
}
