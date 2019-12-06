package api_request_impl

import (
	"bytes"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/network/nw_retry"
	"go.uber.org/zap"
	"net/http"
)

type uploadRequestImpl struct {
	ctx           api_context.Context
	dbxReq        *dbxRequest
	paramString   string
	param         interface{}
	endpoint      string
	url           string
	headers       map[string]string
	method        string
	uploadBytes   []byte
	contentLength int64
}

func (z *uploadRequestImpl) Param(p interface{}) api_request.Request {
	return &uploadRequestImpl{
		ctx:           z.ctx,
		dbxReq:        z.dbxReq,
		paramString:   "",
		param:         p,
		endpoint:      z.endpoint,
		url:           z.url,
		headers:       z.headers,
		method:        z.method,
		uploadBytes:   z.uploadBytes,
		contentLength: z.contentLength,
	}
}

func (z *uploadRequestImpl) Call() (res api_response.Response, err error) {
	return nw_retry.Call(z.ctx, z)
}

func (z *uploadRequestImpl) Endpoint() string {
	return z.endpoint
}

func (z *uploadRequestImpl) ContentLength() int64 {
	return z.contentLength
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

	z.contentLength = int64(len(z.uploadBytes))
	z.headers = make(map[string]string)
	for k := range req.Header {
		z.headers[k] = req.Header.Get(k)
	}

	return req, nil
}
