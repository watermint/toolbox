package sv_activity

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/essentials/format/tjson"
)

type Activity interface {
	All(handler func(event *mo_activity.Event) error) (err error)
	List(handler func(event *mo_activity.Event) error, opts ...ListOpt) error
}

type ListOpt func(opts *ListOpts) *ListOpts
type ListOpts struct {
	StartTime string
	EndTime   string
	Category  string
	AccountId string
}

func StartTime(time string) ListOpt {
	return func(opts *ListOpts) *ListOpts {
		opts.StartTime = time
		return opts
	}
}
func EndTime(time string) ListOpt {
	return func(opts *ListOpts) *ListOpts {
		opts.EndTime = time
		return opts
	}
}
func Category(category string) ListOpt {
	return func(opts *ListOpts) *ListOpts {
		opts.Category = category
		return opts
	}
}
func AccountId(accountId string) ListOpt {
	return func(opts *ListOpts) *ListOpts {
		opts.AccountId = accountId
		return opts
	}
}

func New(ctx dbx_context.Context) Activity {
	return &activityImpl{
		ctx: ctx,
	}
}

type activityImpl struct {
	ctx dbx_context.Context
}

func (z *activityImpl) All(handler func(event *mo_activity.Event) error) (err error) {
	return z.ctx.List("team_log/get_events").
		Continue("team_log/get_events/continue").
		UseHasMore(true).
		ResultTag("events").
		OnEntry(func(entry tjson.Json) error {
			e := &mo_activity.Event{}
			if _, err = entry.Model(e); err != nil {
				return err
			}
			return handler(e)
		}).Call()
}

func (z *activityImpl) List(handler func(event *mo_activity.Event) error, opts ...ListOpt) error {
	los := &ListOpts{}
	for _, o := range opts {
		o(los)
	}
	type TimeRange struct {
		StartTime string `json:"start_time,omitempty"`
		EndTime   string `json:"end_time,omitempty"`
	}
	p := struct {
		AccountId string    `json:"account_id,omitempty"`
		Time      TimeRange `json:"time,omitempty"`
		Category  string    `json:"category,omitempty"`
	}{
		AccountId: los.AccountId,
		Time: TimeRange{
			StartTime: los.StartTime,
			EndTime:   los.EndTime,
		},
		Category: los.Category,
	}

	return z.ctx.List("team_log/get_events").
		Continue("team_log/get_events/continue").
		UseHasMore(true).
		Param(p).
		ResultTag("events").
		OnEntry(func(entry tjson.Json) error {
			e := &mo_activity.Event{}
			if _, err := entry.Model(e); err != nil {
				return err
			}
			return handler(e)
		}).Call()
}
