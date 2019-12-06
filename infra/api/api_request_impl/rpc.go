package api_request_impl

import (
	"bytes"
	"encoding/json"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/network/nw_retry"
	"go.uber.org/zap"
	"net/http"
)

type rpcRequestImpl struct {
	ctx           api_context.Context
	dbxReq        *dbxRequest
	paramString   string
	param         interface{}
	url           string
	endpoint      string
	headers       map[string]string
	method        string
	contentLength int64
}

func (z *rpcRequestImpl) Param(p interface{}) api_request.Request {
	return &rpcRequestImpl{
		ctx:           z.ctx,
		dbxReq:        z.dbxReq,
		paramString:   "",
		param:         p,
		url:           z.url,
		endpoint:      z.endpoint,
		headers:       z.headers,
		method:        z.method,
		contentLength: z.contentLength,
	}
}

func (z *rpcRequestImpl) Call() (res api_response.Response, err error) {
	return nw_retry.Call(z.ctx, z)
}

func (z *rpcRequestImpl) Endpoint() string {
	return z.endpoint
}

func (z *rpcRequestImpl) ContentLength() int64 {
	return z.contentLength
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
	z.contentLength = int64(len(p))
	z.headers = make(map[string]string)
	for k := range req.Header {
		z.headers[k] = req.Header.Get(k)
	}

	return req, nil
}
