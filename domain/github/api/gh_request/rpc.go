package gh_request

import (
	"bytes"
	"encoding/json"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/api/gh_rest"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/app"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"net/http"
)

func NewRpc(ctx gh_context.Context, scope string, token *oauth2.Token, endpoint string, method string) api_request.Request {
	return &rpcImpl{
		ctx:           ctx,
		scope:         scope,
		token:         token,
		endpoint:      endpoint,
		method:        method,
		customHeaders: make(map[string]string),
	}
}

type rpcImpl struct {
	ctx           gh_context.Context
	scope         string
	token         *oauth2.Token
	endpoint      string
	param         interface{}
	method        string
	customHeaders map[string]string

	// mutable variables

	paramString   string
	url           string
	headers       map[string]string
	contentLength int64
}

func (z *rpcImpl) Header(key, value string) api_request.Request {
	headers := make(map[string]string)
	for k, v := range z.customHeaders {
		headers[k] = v
	}
	headers[key] = value
	return &rpcImpl{
		ctx:           z.ctx,
		scope:         z.scope,
		token:         z.token,
		endpoint:      z.endpoint,
		method:        z.method,
		customHeaders: headers,
		param:         z.param,
	}
}

func (z *rpcImpl) ParamString() string {
	return z.paramString
}

func (z *rpcImpl) Param(p interface{}) api_request.Request {
	return &rpcImpl{
		ctx:           z.ctx,
		scope:         z.scope,
		token:         z.token,
		endpoint:      z.endpoint,
		method:        z.method,
		customHeaders: z.customHeaders,
		param:         p,
	}
}

func (z *rpcImpl) Call() (res api_response.Response, err error) {
	return gh_rest.Default().Call(z.ctx, z)
}

func (z *rpcImpl) Endpoint() string {
	return z.endpoint
}

func (z *rpcImpl) Url() string {
	return z.url
}

func (z *rpcImpl) Headers() map[string]string {
	return z.headers
}

func (z *rpcImpl) Method() string {
	return z.method
}

func (z *rpcImpl) ContentLength() int64 {
	return z.contentLength
}

func (z *rpcImpl) Make() (req *http.Request, err error) {
	l := z.ctx.Log().With(zap.String("scope", z.scope), zap.String("endpoint", z.endpoint))

	z.url = "https://api.github.com/" + z.endpoint
	l.Debug("Making request", zap.String("url", z.url))

	p, err := json.Marshal(z.param)
	if err != nil {
		l.Debug("Unable to marshal params", zap.Error(err))
		return nil, err
	}
	z.paramString = string(p)

	req, err = http.NewRequest(z.method, z.url, bytes.NewReader(p))
	if err != nil {
		l.Debug("Unable create request", zap.Error(err))
		return nil, err
	}
	if z.token != nil {
		req.Header.Add(api_request.ReqHeaderAuthorization, "token "+z.token.AccessToken)
	}
	req.Header.Add(api_request.ReqHeaderUserAgent, app.UserAgent())
	req.Header.Add(api_request.ReqHeaderContentType, "application/json")
	for k, v := range z.customHeaders {
		req.Header.Add(k, v)
	}
	z.contentLength = int64(len(p))
	z.headers = make(map[string]string)
	for k := range req.Header {
		z.headers[k] = req.Header.Get(k)
	}
	return req, nil
}
