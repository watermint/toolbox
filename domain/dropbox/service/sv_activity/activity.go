package sv_activity

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/infra/api/api_request"
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
	res := z.ctx.List("team_log/get_events").Call(
		dbx_list.Continue("team_log/get_events/continue"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("events"),
		dbx_list.OnEntry(func(entry tjson.Json) error {
			e := &mo_activity.Event{}
			if err = entry.Model(e); err != nil {
				return err
			}
			return handler(e)
		}),
	)
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
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

	res := z.ctx.List("team_log/get_events", api_request.Param(p)).Call(
		dbx_list.Continue("team_log/get_events/continue"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("events"),
		dbx_list.OnEntry(func(entry tjson.Json) error {
			e := &mo_activity.Event{}
			if err := entry.Model(e); err != nil {
				return err
			}
			return handler(e)
		}),
	)
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}
