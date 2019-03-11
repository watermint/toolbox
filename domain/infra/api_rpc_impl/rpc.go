package api_rpc_impl

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

func New(ec *app.ExecContext,
	endpoint string,
	asMemberId, asAdminId string,
	base api_context.Base,
	token api_auth.Token) api_rpc.Request {

	ri := requestImpl{
		ec:         ec,
		endpoint:   endpoint,
		asMemberId: asMemberId,
		asAdminId:  asAdminId,
		base:       base,
		token:      token,
	}
	return &ri
}

type requestImpl struct {
	ec         *app.ExecContext
	asMemberId string
	asAdminId  string
	base       api_context.Base
	param      interface{}
	token      api_auth.Token
	endpoint   string
	success    func(res api_rpc.Response) error
	failure    func(err error) error
}

func (z *requestImpl) Param(param interface{}) api_rpc.Request {
	z.param = param
	return z
}

func (z *requestImpl) OnSuccess(success func(res api_rpc.Response) error) api_rpc.Request {
	z.success = success
	return z
}

func (z *requestImpl) OnFailure(failure func(err error) error) api_rpc.Request {
	z.failure = failure
	return z
}

func (z *requestImpl) Call() (res api_rpc.Response, err error) {
	rpc := dbx_rpc.RpcRequest{
		Endpoint:   z.endpoint,
		Param:      z.param,
		AsMemberId: z.asMemberId,
		AsAdminId:  z.asAdminId,
		//PathRoot: z.base, TODO: incompatible
	}
	ctx := dbx_api.NewContext(z.ec, "api_rpc_impl", z.token.Token())
	dbxRes, err := rpc.Call(ctx)
	if err != nil {
		if z.failure != nil {
			return newFailureResponse(err), z.failure(err)
		}
		return newFailureResponse(err), err
	}
	apiRes := newSuccessResponse(dbxRes)
	if z.success != nil {
		return apiRes, z.success(apiRes)
	}
	return apiRes, nil
}

func newFailureResponse(resErr error) api_rpc.Response {
	return &responseImpl{
		resErr: resErr,
	}
}

func newSuccessResponse(dbxRes *dbx_rpc.RpcResponse) api_rpc.Response {
	return &responseImpl{
		dbxRes: dbxRes,
	}
}

type responseImpl struct {
	resErr error
	dbxRes *dbx_rpc.RpcResponse
}

func (z *responseImpl) Error() error {
	return z.resErr
}

func (z *responseImpl) StatusCode() int {
	if z.dbxRes != nil {
		return z.dbxRes.StatusCode
	}
	return 0
}

func (z *responseImpl) Body() (body string, err error) {
	if z.dbxRes != nil {
		return z.dbxRes.Body, nil
	}
	return "", errors.New("no body")
}

func (z *responseImpl) Json() (res gjson.Result, err error) {
	body, err := z.Body()
	if err != nil {
		return gjson.Parse(`{}`), err
	}
	if !gjson.Valid(body) {
		return gjson.Parse(`{}`), errors.New("not a json data")
	}
	return gjson.Parse(body), nil
}

func (z *responseImpl) Model(v interface{}) error {
	body, err := z.Body()
	if err != nil {
		return err
	}
	return api_parser.ParseModelString(v, body)
}

func (z *responseImpl) ModelWithPath(v interface{}, path string) error {
	body, err := z.Body()
	if err != nil {
		return err
	}
	if !gjson.Valid(body) {
		return errors.New("not a json data")
	}
	p := gjson.Get(body, path)
	if !p.Exists() {
		return errors.New("data not found for path")
	}
	return api_parser.ParseModel(v, p)
}
