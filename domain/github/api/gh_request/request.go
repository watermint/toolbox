package gh_request

import (
	"bytes"
	"encoding/json"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/api/gh_rest"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"net/http"
)

func New(ctx gh_context.Context, scope string, token *oauth2.Token, endpoint string, method string) api_request.Request {
	return &rpcImpl{
		ctx:      ctx,
		scope:    scope,
		token:    token,
		endpoint: endpoint,
		method:   method,
	}
}

type rpcImpl struct {
	ctx      gh_context.Context
	scope    string
	token    *oauth2.Token
	endpoint string
	param    interface{}
	method   string

	// mutable variables

	paramString   string
	url           string
	headers       map[string]string
	contentLength int64
}

func (z *rpcImpl) ParamString() string {
	return z.paramString
}

func (z *rpcImpl) Param(p interface{}) api_request.Request {
	return &rpcImpl{
		ctx:      z.ctx,
		scope:    z.scope,
		token:    z.token,
		endpoint: z.endpoint,
		method:   z.method,
		param:    p,
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

	req.Header.Add(api_request.ReqHeaderContentType, "application/json")
	z.contentLength = int64(len(p))
	z.headers = make(map[string]string)
	for k := range req.Header {
		z.headers[k] = req.Header.Get(k)
	}
	return req, nil
}
