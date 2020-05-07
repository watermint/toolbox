package dbx_list

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/http/es_response"
)

type List interface {
	Call(opts ...ListOpt) dbx_response.Response
}

type ListOpts struct {
	ContinueEndpoint string
	UseHasMore       bool
	ResultTag        string
	onResponse       func(res es_response.Response) error
	onEntry          func(entry es_json.Json) error
	onLastCursor     func(cursor string)
}

func (z ListOpts) HasOnEntry() bool {
	return z.onEntry != nil
}
func (z ListOpts) HasOnResponse() bool {
	return z.onResponse != nil
}
func (z ListOpts) HasLastCursor() bool {
	return z.onLastCursor != nil
}

func (z ListOpts) OnResponse(res es_response.Response) error {
	if z.onResponse != nil {
		return z.onResponse(res)
	}
	return nil
}
func (z ListOpts) OnEntry(entry es_json.Json) error {
	if z.onEntry != nil {
		return z.onEntry(entry)
	}
	return nil
}
func (z ListOpts) OnLastCursor(cursor string) {
	if z.onLastCursor != nil {
		z.onLastCursor(cursor)
	}
}

type ListOpt func(o ListOpts) ListOpts

func Continue(endpoint string) ListOpt {
	return func(o ListOpts) ListOpts {
		o.ContinueEndpoint = endpoint
		return o
	}
}
func UseHasMore() ListOpt {
	return func(o ListOpts) ListOpts {
		o.UseHasMore = true
		return o
	}
}
func ResultTag(tag string) ListOpt {
	return func(o ListOpts) ListOpts {
		o.ResultTag = tag
		return o
	}
}
func OnResponse(f func(res es_response.Response) error) ListOpt {
	return func(o ListOpts) ListOpts {
		o.onResponse = f
		return o
	}
}
func OnEntry(f func(entry es_json.Json) error) ListOpt {
	return func(o ListOpts) ListOpts {
		o.onEntry = f
		return o
	}
}
func OnLastCursor(f func(cursor string)) ListOpt {
	return func(o ListOpts) ListOpts {
		o.onLastCursor = f
		return o
	}
}

func Combined(opts []ListOpt) ListOpts {
	lo := ListOpts{}
	for _, o := range opts {
		lo = o(lo)
	}
	return lo
}
