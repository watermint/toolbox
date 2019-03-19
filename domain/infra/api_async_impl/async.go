package api_async_impl

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_async"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"github.com/watermint/toolbox/domain/infra/api_rpc_impl"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

func New(ec *app.ExecContext, ctx api_context.Context, endpoint string, asMemberId, asAdminId string, base api_context.Base, token api_auth.Token) api_async.Async {
	return &asyncImpl{
		ec:              ec,
		ctx:             ctx,
		requestEndpoint: endpoint,
		asMemberId:      asMemberId,
		asAdminId:       asAdminId,
		base:            base,
	}
}

type asyncImpl struct {
	ctx             api_context.Context
	ec              *app.ExecContext
	asMemberId      string
	asAdminId       string
	base            api_context.Base
	param           interface{}
	token           api_auth.Token
	pollInterval    int
	requestEndpoint string
	statusEndpoint  string
	success         func(res api_async.Response) error
	failure         func(err error) error
}

func (z *asyncImpl) Param(p interface{}) api_async.Async {
	z.param = p
	return z
}

func (z *asyncImpl) Status(endpoint string) api_async.Async {
	z.statusEndpoint = endpoint
	return z
}

func (z *asyncImpl) PollInterval(second int) api_async.Async {
	z.pollInterval = second
	return z
}

func (z *asyncImpl) OnSuccess(success func(res api_async.Response) error) api_async.Async {
	z.success = success
	return z
}

func (z *asyncImpl) OnFailure(failure func(err error) error) api_async.Async {
	z.failure = failure
	return z
}

func (z *asyncImpl) Call() (res api_async.Response, resErr error) {
	rpcReq := z.ctx.Request(z.requestEndpoint).Param(z.param)
	rpcRes, err := rpcReq.Call()
	if err != nil {
		if z.failure != nil {
			return newFailureResponse(err), z.failure(err)
		}
		return newFailureResponse(err), err
	}
	switch ri := rpcRes.(type) {
	case *api_rpc_impl.ResponseImpl:
		switch qi := rpcReq.(type) {
		case *api_rpc_impl.RequestImpl:
			as := dbx_rpc.AsyncStatus{
				Endpoint:   z.statusEndpoint,
				AsMemberId: z.asMemberId,
				AsAdminId:  z.asAdminId,
				OnError: func(err error) bool {
					res = newFailureResponse(err)
					resErr = err

					if z.failure != nil {
						resErr = z.failure(err)
					}
					return resErr != nil
				},
				OnComplete: func(complete gjson.Result) bool {
					res = newSuccessResponse(ri, complete)
					if z.success != nil {
						resErr = z.success(res)
					}
					return true
				},
			}

			as.Poll(qi.DbxApiContext(), ri.DbxRpcRes())
			return res, resErr
		}
	}
	panic("invalid response type")
}

func newSuccessResponse(resImpl *api_rpc_impl.ResponseImpl, complete gjson.Result) *responseImpl {
	return &responseImpl{
		resImpl:        resImpl,
		completeExists: true,
		complete:       complete,
	}
}

func newFailureResponse(err error) *responseImpl {
	return &responseImpl{
		resErr: err,
	}
}

type responseImpl struct {
	resErr         error
	resImpl        *api_rpc_impl.ResponseImpl
	complete       gjson.Result
	completeExists bool
}

func (z *responseImpl) Error() error {
	return z.resErr
}

func (z *responseImpl) Json() (res gjson.Result, err error) {
	if !z.completeExists {
		return gjson.Parse("{}"), errors.New("no result")
	}
	return z.complete, nil
}

func (z *responseImpl) Model(v interface{}) error {
	if !z.completeExists {
		return errors.New("no result")
	}
	return api_parser.ParseModel(v, z.complete)
}

func (z *responseImpl) ModelWithPath(v interface{}, path string) error {
	if !z.completeExists {
		return errors.New("no result")
	}
	return api_parser.ParseModel(v, z.complete.Get(path))
}
