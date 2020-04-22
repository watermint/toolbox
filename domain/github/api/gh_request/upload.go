package gh_request

import (
	"github.com/google/go-querystring/query"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/api/gh_rest"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"net/http"
)

func NewUpload(ctx gh_context.Context, scope string, token *oauth2.Token, endpoint string, method string, content ut_io.ReadRewinder) api_request.Request {
	return &uploadImpl{
		ctx:           ctx,
		scope:         scope,
		token:         token,
		endpoint:      endpoint,
		method:        method,
		content:       content,
		customHeaders: make(map[string]string),
	}
}

type uploadImpl struct {
	ctx           gh_context.Context
	scope         string
	token         *oauth2.Token
	endpoint      string
	param         interface{}
	method        string
	content       ut_io.ReadRewinder
	customHeaders map[string]string

	// mutable variables

	paramString   string
	url           string
	headers       map[string]string
	contentLength int64
}

func (z *uploadImpl) Header(key, value string) api_request.Request {
	headers := make(map[string]string)
	for k, v := range z.customHeaders {
		headers[k] = v
	}
	headers[key] = value
	return &uploadImpl{
		ctx:           z.ctx,
		scope:         z.scope,
		token:         z.token,
		endpoint:      z.endpoint,
		method:        z.method,
		customHeaders: headers,
		param:         z.param,
		content:       z.content,
	}
}

func (z *uploadImpl) ParamString() string {
	return z.paramString
}

func (z *uploadImpl) Param(p interface{}) api_request.Request {
	return &uploadImpl{
		ctx:      z.ctx,
		scope:    z.scope,
		token:    z.token,
		endpoint: z.endpoint,
		method:   z.method,
		content:  z.content,
		param:    p,
	}
}

func (z *uploadImpl) Call() (res api_response.Response, err error) {
	return gh_rest.Default(z.ctx.Feature()).Call(z.ctx, z)
}

func (z *uploadImpl) Endpoint() string {
	return z.endpoint
}

func (z *uploadImpl) Url() string {
	return z.url
}

func (z *uploadImpl) Headers() map[string]string {
	return z.headers
}

func (z *uploadImpl) Method() string {
	return z.method
}

func (z *uploadImpl) ContentLength() int64 {
	return z.contentLength
}

func (z *uploadImpl) Make() (req *http.Request, err error) {
	l := z.ctx.Log().With(zap.String("scope", z.scope), zap.String("endpoint", z.endpoint))

	qs, err := query.Values(z.param)
	if err != nil {
		l.Debug("Unable to marshal params", zap.Error(err))
		return nil, err
	}
	z.paramString = qs.Encode()
	z.url = "https://uploads.github.com/" + z.endpoint + "?" + z.paramString

	l.Debug("Making request", zap.String("url", z.url))

	req, err = http.NewRequest(z.method, z.url, nw_bandwidth.WrapReader(z.content))
	if err != nil {
		l.Debug("Unable create request", zap.Error(err))
		return nil, err
	}

	// content length
	z.contentLength = z.content.Length()
	req.ContentLength = z.contentLength

	// content type
	customContentType := false
	for k := range z.customHeaders {
		if k == api_request.ReqHeaderContentType {
			customContentType = true
			break
		}
	}
	if !customContentType {
		req.Header.Add(api_request.ReqHeaderContentType, "application/json")
	}

	// custom headers
	for k, v := range z.customHeaders {
		req.Header.Add(k, v)
	}

	// auth token
	if z.token != nil {
		req.Header.Add(api_request.ReqHeaderAuthorization, "token "+z.token.AccessToken)
	}

	// user agent
	req.Header.Add(api_request.ReqHeaderUserAgent, app.UserAgent())

	// save headers for logging
	z.headers = make(map[string]string)
	for k := range req.Header {
		z.headers[k] = req.Header.Get(k)
	}
	return req, nil
}
