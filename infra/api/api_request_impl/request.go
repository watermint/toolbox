package api_request_impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/network/nw_retry"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	RpcEndpoint     = "api.dropboxapi.com"
	ContentEndpoint = "content.dropboxapi.com"
)

func RpcRequestUrl(endpoint string) string {
	return fmt.Sprintf("https://%s/2/%s", RpcEndpoint, endpoint)
}

func ContentRequestUrl(endpoint string) string {
	return fmt.Sprintf("https://%s/2/%s", ContentEndpoint, endpoint)
}

func NewPpcRequest(ctx api_context.Context,
	endpoint string,
	asMemberId, asAdminId string,
	base api_context.PathRoot,
	token api_auth.TokenContainer) api_request.Request {

	req := &rpcRequestImpl{
		ctx:      ctx,
		endpoint: endpoint,
		dbxReq: &dbxRequest{
			asMemberId: asMemberId,
			asAdminId:  asAdminId,
			base:       base,
			token:      token,
		},
	}
	return req
}

func NewUploadRequest(ctx api_context.Context,
	endpoint string,
	content io.Reader,
	asMemberId, asAdminId string,
	base api_context.PathRoot,
	token api_auth.TokenContainer) api_request.Request {

	l := app_root.Log()
	upload, err := ioutil.ReadAll(content)
	if err != nil {
		l.Debug("Unable to read", zap.Error(err))
		return &badRequest{err: err}
	}

	req := &uploadRequestImpl{
		ctx:      ctx,
		endpoint: endpoint,
		dbxReq: &dbxRequest{
			asMemberId: asMemberId,
			asAdminId:  asAdminId,
			base:       base,
			token:      token,
		},
		uploadBytes: upload,
	}
	return req
}

type badRequest struct {
	err error
}

func (z *badRequest) ParamString() string {
	return ""
}

func (z *badRequest) Param(p interface{}) api_request.Request {
	return &badRequest{
		err: z.err,
	}
}

func (z *badRequest) Call() (res api_response.Response, err error) {
	return nil, z.err
}

func (z *badRequest) Endpoint() string {
	return ""
}

func (z *badRequest) Url() string {
	return ""
}

func (z *badRequest) Headers() map[string]string {
	return map[string]string{}
}

func (z *badRequest) Method() string {
	return ""
}

func (z *badRequest) Make() (req *http.Request, err error) {
	return nil, z.err
}

type dbxRequest struct {
	asMemberId string
	asAdminId  string
	base       api_context.PathRoot
	token      api_auth.TokenContainer
}

func (z *dbxRequest) decorate(req *http.Request) (r *http.Request, err error) {
	l := app_root.Log()

	if z.token.TokenType != api_auth.DropboxTokenNoAuth {
		req.Header.Add(api_request.ReqHeaderAuthorization, "Bearer "+z.token.Token)
	}
	if z.asAdminId != "" {
		req.Header.Add(api_request.ReqHeaderSelectAdmin, z.asAdminId)
	}
	if z.asMemberId != "" {
		req.Header.Add(api_request.ReqHeaderSelectUser, z.asMemberId)
	}
	if z.base != nil {
		pr, err := json.Marshal(z.base)
		if err != nil {
			l.Debug("unable to marshal path root", zap.Error(err))
			return nil, err
		}
		req.Header.Add(api_request.ReqHeaderPathRoot, string(pr))
	}
	return req, nil
}

type rpcRequestImpl struct {
	ctx         api_context.Context
	dbxReq      *dbxRequest
	paramString string
	param       interface{}
	url         string
	endpoint    string
	headers     map[string]string
	method      string
}

func (z *rpcRequestImpl) Param(p interface{}) api_request.Request {
	return &rpcRequestImpl{
		ctx:         z.ctx,
		dbxReq:      z.dbxReq,
		paramString: "",
		param:       p,
		url:         z.url,
		endpoint:    z.endpoint,
		headers:     z.headers,
		method:      z.method,
	}
}

func (z *rpcRequestImpl) Call() (res api_response.Response, err error) {
	return nw_retry.Call(z.ctx, z)
}

func (z *rpcRequestImpl) Endpoint() string {
	return z.endpoint
}

func (z *rpcRequestImpl) ParamString() string {
	return z.paramString
}

func (z *rpcRequestImpl) Url() string {
	return z.url
}

func (z *rpcRequestImpl) Headers() map[string]string {
	return z.headers
}

func (z *rpcRequestImpl) Method() string {
	return z.method
}

func (z *rpcRequestImpl) Make() (req *http.Request, err error) {
	l := z.ctx.Log()

	z.url = RpcRequestUrl(z.endpoint)

	// param
	p, err := json.Marshal(z.param)
	if err != nil {
		l.Debug("Unable to marshal params", zap.Error(err))
		return nil, err
	}
	z.paramString = string(p)
	z.method = "POST"

	req, err = http.NewRequest(z.method, z.url, bytes.NewReader(p))
	if err != nil {
		l.Debug("Unable create request", zap.Error(err))
		return nil, err
	}
	if _, err := z.dbxReq.decorate(req); err != nil {
		return nil, err
	}
	req.Header.Add(api_request.ReqHeaderContentType, "application/json")
	z.headers = make(map[string]string)
	for k := range req.Header {
		z.headers[k] = req.Header.Get(k)
	}

	return req, nil
}

type uploadRequestImpl struct {
	ctx         api_context.Context
	dbxReq      *dbxRequest
	paramString string
	param       interface{}
	endpoint    string
	url         string
	headers     map[string]string
	method      string
	uploadBytes []byte
}

func (z *uploadRequestImpl) Param(p interface{}) api_request.Request {
	return &uploadRequestImpl{
		ctx:         z.ctx,
		dbxReq:      z.dbxReq,
		paramString: "",
		param:       p,
		endpoint:    z.endpoint,
		url:         z.url,
		headers:     z.headers,
		method:      z.method,
		uploadBytes: z.uploadBytes,
	}
}

func (z *uploadRequestImpl) Call() (res api_response.Response, err error) {
	return nw_retry.Call(z.ctx, z)
}

func (z *uploadRequestImpl) Endpoint() string {
	return z.endpoint
}

func (z *uploadRequestImpl) ParamString() string {
	return z.paramString
}

func (z *uploadRequestImpl) Url() string {
	return z.url
}

func (z *uploadRequestImpl) Headers() map[string]string {
	return z.headers
}

func (z *uploadRequestImpl) Method() string {
	return z.method
}

func (z *uploadRequestImpl) Make() (req *http.Request, err error) {
	l := z.ctx.Log()
	z.url = ContentRequestUrl(z.endpoint)
	z.method = "POST"

	req, err = http.NewRequest(z.method, z.url, nw_bandwidth.WrapReader(bytes.NewReader(z.uploadBytes)))
	if err != nil {
		l.Debug("Unable create request", zap.Error(err))
		return nil, err
	}
	if _, err := z.dbxReq.decorate(req); err != nil {
		return nil, err
	}
	z.paramString, err = api_util.HeaderSafeJson(z.param)
	if err != nil {
		l.Debug("Unable to encode json", zap.Error(err))
		return nil, err
	}

	req.Header.Add(api_request.ReqHeaderContentType, "application/octet-stream")
	req.Header.Add(api_request.ReqHeaderArg, z.paramString)

	z.headers = make(map[string]string)
	for k := range req.Header {
		z.headers[k] = req.Header.Get(k)
	}

	return req, nil
}
