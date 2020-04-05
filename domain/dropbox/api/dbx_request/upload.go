package dbx_request

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/network/nw_retry"
	"github.com/watermint/toolbox/infra/util/ut_io"
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
	content       ut_io.ReadRewinder
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
		content:       z.content,
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

	if err = z.content.Rewind(); err != nil {
		l.Debug("Unable to rewind", zap.Error(err))
		return nil, err
	}
	req, err = http.NewRequest(z.method, z.url, nw_bandwidth.WrapReader(z.content))
	if err != nil {
		l.Debug("Unable create request", zap.Error(err))
		return nil, err
	}
	if _, err := z.dbxReq.decorate(req); err != nil {
		return nil, err
	}
	z.paramString, err = dbx_util.HeaderSafeJson(z.param)
	if err != nil {
		l.Debug("Unable to encode json", zap.Error(err))
		return nil, err
	}

	req.Header.Add(api_request.ReqHeaderUserAgent, app.UserAgent())
	req.Header.Add(api_request.ReqHeaderContentType, "application/octet-stream")
	req.Header.Add(api_request.ReqHeaderDropboxApiArg, z.paramString)

	z.contentLength = z.content.Length()
	z.headers = make(map[string]string)
	for k := range req.Header {
		z.headers[k] = req.Header.Get(k)
	}

	return req, nil
}
