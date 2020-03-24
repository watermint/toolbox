package api_request_impl

import (
	"bytes"
	"encoding/json"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/network/nw_retry"
	"go.uber.org/zap"
	"net/http"
)

type downloadRequestImpl struct {
	ctx           api_context.Context
	dbxReq        *dbxRequest
	paramString   string
	param         interface{}
	endpoint      string
	url           string
	headers       map[string]string
	method        string
	contentLength int64
}

func (z *downloadRequestImpl) ParamString() string {
	return z.paramString
}

func (z *downloadRequestImpl) ContentLength() int64 {
	return z.contentLength
}

func (z *downloadRequestImpl) Param(p interface{}) api_request.Request {
	return &downloadRequestImpl{
		ctx:           z.ctx,
		dbxReq:        z.dbxReq,
		paramString:   "",
		param:         p,
		endpoint:      z.endpoint,
		url:           z.url,
		headers:       z.headers,
		method:        z.method,
		contentLength: z.contentLength,
	}
}

func (z *downloadRequestImpl) Call() (res api_response.Response, err error) {
	return nw_retry.Call(z.ctx, z)
}

func (z *downloadRequestImpl) Endpoint() string {
	return z.endpoint
}

func (z *downloadRequestImpl) Url() string {
	return z.url
}

func (z *downloadRequestImpl) Headers() map[string]string {
	return z.headers
}

func (z *downloadRequestImpl) Method() string {
	return z.method
}

func (z *downloadRequestImpl) Make() (req *http.Request, err error) {
	l := z.ctx.Log()
	z.url = ContentRequestUrl(z.endpoint)

	// param
	p, err := json.Marshal(z.param)
	if err != nil {
		l.Debug("Unable to marshal params", zap.Error(err))
		return nil, err
	}
	z.paramString = string(p)
	z.method = "POST"

	req, err = http.NewRequest(z.method, z.url, bytes.NewReader([]byte{}))
	if err != nil {
		l.Debug("Unable create request", zap.Error(err))
		return nil, err
	}
	if _, err := z.dbxReq.decorate(req); err != nil {
		return nil, err
	}
	req.Header.Add(api_request.ReqHeaderUserAgent, app.UserAgent())
	req.Header.Add(api_request.ReqHeaderContentType, "application/octet-stream")
	req.Header.Add(api_request.ReqHeaderDropboxApiArg, z.paramString)
	z.contentLength = 0

	z.headers = make(map[string]string)
	for k := range req.Header {
		z.headers[k] = req.Header.Get(k)
	}

	return req, nil
}
