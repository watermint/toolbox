package dbx_list_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response_impl"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_request"
	"go.uber.org/zap"
)

var (
	ErrorNoResult = errors.New("no result")
)

func New(ctx dbx_context.Context, endpoint string, reqData []api_request.RequestDatum) dbx_list.List {
	return &listImpl{
		ctx:         ctx,
		reqData:     reqData,
		reqEndpoint: endpoint,
	}
}

type listImpl struct {
	ctx         dbx_context.Context
	reqData     []api_request.RequestDatum
	reqEndpoint string
}

func (z listImpl) log(lo dbx_list.ListOpts) *zap.Logger {
	return z.ctx.Log().With(
		zap.String("reqEndpoint", z.reqEndpoint),
		zap.String("contEndpoint", lo.ContinueEndpoint),
	)
}

func (z *listImpl) handleResponse(lo dbx_list.ListOpts, res dbx_response.Response) dbx_response.Response {
	l := z.log(lo)

	if err, fail := res.Failure(); fail {
		l.Debug("error response", zap.Error(err))
		return res
	}

	l.Debug("on response")
	if err := lo.OnResponse(res); err != nil {
		return dbx_response_impl.NewAbort(res, err)
	}

	l.Debug("handle entry")
	if err := z.handleEntry(lo, res); err != nil {
		return dbx_response_impl.NewAbort(res, err)
	}

	l.Debug("determine continue")
	if cont, cursor := z.isContinue(lo, res); cont {

		l.Debug("continue")
		return z.listContinue(lo, cursor)
	} else {
		lo.OnLastCursor(cursor)
		return dbx_response_impl.New(res)
	}
}

func (z listImpl) handleEntry(lo dbx_list.ListOpts, res dbx_response.Response) error {
	l := z.log(lo)
	if err, fail := res.Failure(); fail {
		return err
	}

	j := res.Result()

	if results, found := j.FindArray(lo.ResultTag); !found {
		l.Debug("No result found", zap.ByteString("response", j.Raw()))
		return ErrorNoResult
	} else {
		for _, e := range results {
			if err := lo.OnEntry(e); err != nil {
				l.Debug("handler returned abort", zap.Error(err))
				return err
			}
		}
		return nil
	}
}

func (z listImpl) isContinueHasMore(lo dbx_list.ListOpts, j tjson.Json) (cont bool, cursor string) {
	l := z.log(lo)
	if hasMore, e := j.FindBool("has_more"); !hasMore {
		l.Debug("no more results; has_more == false",
			zap.Bool("e", e),
			zap.Bool("hasMore", hasMore))
		return false, ""
	}
	return z.isContinueCursor(lo, j)
}

func (z listImpl) isContinueCursor(lo dbx_list.ListOpts, j tjson.Json) (cont bool, cursor string) {
	l := z.log(lo)
	if cursor, found := j.FindString("cursor"); found {
		l.Debug("cursor found", zap.String("cursor", cursor))
		return true, cursor
	} else {
		l.Debug("has_more returned true, but no cursor found in the body")
		return false, ""
	}
}

func (z listImpl) isContinue(lo dbx_list.ListOpts, res response.Response) (cont bool, cursor string) {
	l := z.log(lo)
	j, err := res.Success().AsJson()
	if err != nil {
		return false, ""
	}

	if lo.UseHasMore {
		l.Debug("use has_more")
		return z.isContinueHasMore(lo, j)
	} else {
		l.Debug("no has_more defined for this api")
		return z.isContinueCursor(lo, j)
	}
}

func (z listImpl) list(lo dbx_list.ListOpts) dbx_response.Response {
	res := z.ctx.Post(z.reqEndpoint, z.reqData...)
	return z.handleResponse(lo, res)
}

func (z listImpl) listContinue(lo dbx_list.ListOpts, cursor string) dbx_response.Response {
	p := struct {
		Cursor string `json:"cursor"`
	}{
		Cursor: cursor,
	}
	res := z.ctx.Post(lo.ContinueEndpoint, api_request.Param(p))

	return z.handleResponse(lo, res)
}

func (z listImpl) Call(opts ...dbx_list.ListOpt) dbx_response.Response {
	lo := dbx_list.Combined(opts)
	return z.list(lo)
}
