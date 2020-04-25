package dbx_list

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_list"
	"go.uber.org/zap"
)

var (
	ErrorNoResult = errors.New("no result")
)

func New(ctx dbx_context.Context, endpoint string, asMemberId, asAdminId string, base dbx_context.PathRoot) api_list.List {
	return &listImpl{
		ctx:             ctx,
		requestEndpoint: endpoint,
		asMemberId:      asMemberId,
		asAdminId:       asAdminId,
		base:            base,
	}
}

type listImpl struct {
	ctx              dbx_context.Context
	asMemberId       string
	asAdminId        string
	base             dbx_context.PathRoot
	param            interface{}
	token            api_auth.Context
	useHasMore       bool
	resultTag        string
	requestEndpoint  string
	continueEndpoint string
	onEntry          func(res tjson.Json) error
	onResponse       func(res response.Response) error
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

func (z *listImpl) OnResponse(response func(res response.Response) error) api_list.List {
	z.onResponse = response
	return z
}

func (z *listImpl) OnEntry(entry func(entry tjson.Json) error) api_list.List {
	z.onEntry = entry
	return z
}

func (z *listImpl) OnLastCursor(f func(cursor string)) api_list.List {
	z.onLastCursor = f
	return z
}

func (z *listImpl) handleResponse(endpoint string, res response.Response, err error) error {
	l := z.ctx.Log().With(zap.String("endpoint", endpoint))

	if err != nil {
		return err
	}

	if res == nil {
		l.Warn("Response is null")
		return errors.New("response is null")
	}

	if z.onResponse != nil {
		if err = z.onResponse(res); err != nil {
			l.Debug("OnResponseBody returned abort", zap.Error(err))
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

func (z *listImpl) handleEntry(res response.Response) error {
	if z.onEntry == nil {
		return nil
	}

	l := z.ctx.Log().With(zap.String("endpoint", z.requestEndpoint), zap.String("result_tag", z.resultTag))
	if err := dbx_error.IsApiError(res); err != nil {
		return err
	}
	j := res.Success().Json()

	if resultsElem, found := j.Find(z.resultTag); !found {
		l.Debug("No result found", zap.ByteString("response", j.Raw()))
		return ErrorNoResult
	} else if results, found := resultsElem.Array(); !found {
		l.Debug("Result was not an array", zap.ByteString("response", j.Raw()))
		return ErrorNoResult
	} else {
		for _, e := range results {
			if err := z.onEntry(e); err != nil {
				l.Debug("handler returned abort", zap.Error(err))
				return err
			}
		}
		return nil
	}
}

func (z listImpl) isContinueHasMore(j tjson.Json) (cont bool, cursor string) {
	l := z.ctx.Log().With(zap.String("endpoint", z.continueEndpoint))
	if hasMore, e := j.FindBool("has_more"); !hasMore {
		l.Debug("no more results; has_more == false",
			zap.Bool("e", e),
			zap.Bool("hasMore", hasMore))
		return false, ""
	}
	return z.isContinueCursor(j)
}

func (z listImpl) isContinueCursor(j tjson.Json) (cont bool, cursor string) {
	l := z.ctx.Log().With(zap.String("endpoint", z.continueEndpoint))
	if cursor, found := j.FindString("cursor"); found {
		l.Debug("cursor found", zap.String("cursor", cursor))
		return true, cursor
	} else {
		l.Debug("has_more returned true, but no cursor found in the body")
		return false, ""
	}
}

func (z *listImpl) isContinue(res response.Response) (cont bool, cursor string) {
	l := z.ctx.Log().With(zap.String("endpoint", z.continueEndpoint))
	j, err := res.Success().AsJson()
	if err != nil {
		return false, ""
	}

	if z.useHasMore {
		l.Debug("use has_more")
		return z.isContinueHasMore(j)
	} else {
		l.Debug("no has_more defined for this api")
		return z.isContinueCursor(j)
	}
}

func (z *listImpl) list() error {
	res, err := z.ctx.Post(z.requestEndpoint).Param(z.param).Call()
	return z.handleResponse(z.requestEndpoint, res, err)
}

func (z *listImpl) listContinue(cursor string) error {
	p := struct {
		Cursor string `json:"cursor"`
	}{
		Cursor: cursor,
	}
	res, err := z.ctx.Post(z.continueEndpoint).Param(p).Call()

	return z.handleResponse(z.continueEndpoint, res, err)
}

func (z *listImpl) Call() (err error) {
	return z.list()
}
