package dbx_list

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_response"
	"go.uber.org/zap"
)

func New(ctx api_context.Context, endpoint string, asMemberId, asAdminId string, base api_context.PathRoot) api_list.List {
	return &listImpl{
		ctx:             ctx,
		requestEndpoint: endpoint,
		asMemberId:      asMemberId,
		asAdminId:       asAdminId,
		base:            base,
	}
}

type listImpl struct {
	ctx              api_context.Context
	asMemberId       string
	asAdminId        string
	base             api_context.PathRoot
	param            interface{}
	token            api_auth.TokenContainer
	useHasMore       bool
	resultTag        string
	requestEndpoint  string
	continueEndpoint string
	onEntry          func(res api_list.ListEntry) error
	onResponse       func(res api_response.Response) error
	onFailure        func(err error) error
	onLastCursor     func(cursor string)
}

func (z *listImpl) Param(param interface{}) api_list.List {
	z.param = param
	return z
}

func (z *listImpl) Continue(endpoint string) api_list.List {
	z.continueEndpoint = endpoint
	return z
}

func (z *listImpl) UseHasMore(use bool) api_list.List {
	z.useHasMore = use
	return z
}

func (z *listImpl) ResultTag(tag string) api_list.List {
	z.resultTag = tag
	return z
}

func (z *listImpl) OnFailure(failure func(err error) error) api_list.List {
	z.onFailure = failure
	return z
}

func (z *listImpl) OnResponse(response func(res api_response.Response) error) api_list.List {
	z.onResponse = response
	return z
}

func (z *listImpl) OnEntry(entry func(entry api_list.ListEntry) error) api_list.List {
	z.onEntry = entry
	return z
}

func (z *listImpl) OnLastCursor(f func(cursor string)) api_list.List {
	z.onLastCursor = f
	return z
}

func (z *listImpl) handleResponse(endpoint string, res api_response.Response, err error) error {
	log := z.ctx.Log().With(zap.String("endpoint", endpoint))

	if err != nil {
		if z.onFailure != nil {
			return z.onFailure(err)
		}
		return err
	}

	if res == nil {
		log.Warn("Response is null")
		return errors.New("response is null")
	}

	if z.onResponse != nil {
		if err = z.onResponse(res); err != nil {
			log.Debug("OnResponseBody returned abort", zap.Error(err))
			return err
		}
	}

	if z.onEntry != nil {
		if err = z.handleEntry(res); err != nil {
			return err
		}
	}

	if cont, cursor := z.isContinue(res); cont {
		if err = z.listContinue(cursor); err != nil {
			return err
		}
	} else if z.onLastCursor != nil {
		z.onLastCursor(cursor)
	}
	return nil
}

func (z *listImpl) handleEntry(res api_response.Response) error {
	if z.onEntry == nil {
		return nil
	}

	log := z.ctx.Log().With(zap.String("endpoint", z.requestEndpoint), zap.String("result_tag", z.resultTag))
	j, err := res.Json()
	if err != nil {
		return err
	}

	results := j.Get(z.resultTag)
	if !results.Exists() {
		log.Debug("No result found")
		return errors.New("no result found")
	}

	if !results.IsArray() {
		log.Debug("result was not an array")
		return errors.New("result was not an array")
	}

	for _, e := range results.Array() {
		le := &listEntryImpl{
			entry: e,
		}
		if err = z.onEntry(le); err != nil {
			log.Debug("handler returned abort")
			return err
		}
	}
	return nil
}

func (z *listImpl) isContinue(res api_response.Response) (cont bool, cursor string) {
	log := z.ctx.Log().With(zap.String("endpoint", z.continueEndpoint))
	j, err := res.Json()
	if err != nil {
		return false, ""
	}

	if z.useHasMore {
		cursor = j.Get("cursor").String()
		if j.Get("has_more").Bool() {
			if cursor != "" {
				return true, cursor
			}
			log.Debug("has_more returned true, but no cursor found in the body")
			return false, ""
		} else {
			return false, cursor
		}
	}

	cursor = j.Get("cursor").String()
	if cursor != "" {
		return true, cursor
	}
	return false, cursor
}

func (z *listImpl) list() error {
	res, err := z.ctx.Rpc(z.requestEndpoint).Param(z.param).Call()
	return z.handleResponse(z.requestEndpoint, res, err)
}

func (z *listImpl) listContinue(cursor string) error {
	p := struct {
		Cursor string `json:"cursor"`
	}{
		Cursor: cursor,
	}
	res, err := z.ctx.Rpc(z.continueEndpoint).Param(p).Call()

	return z.handleResponse(z.continueEndpoint, res, err)
}

func (z *listImpl) Call() (err error) {
	return z.list()
}
