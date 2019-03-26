package api_rpc_impl

import (
	"bytes"
	"encoding/json"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"go.uber.org/zap"
	"net/http"
)

type requestImpl struct {
	ctx     api_context.Context
	param   []byte
	url     string
	headers map[string]string
	method  string
}

func (z *requestImpl) log() *zap.Logger {
	return z.ctx.Log()
}

func (z *requestImpl) Param() string {
	return string(z.param)
}

func (z *requestImpl) Url() string {
	return z.url
}

func (z *requestImpl) Headers() map[string]string {
	return z.headers
}

func (z *requestImpl) Method() string {
	return z.method
}

func (z *requestImpl) Request() (req *http.Request, err error) {
	req, err = http.NewRequest(z.method, z.url, bytes.NewReader(z.param))
	if err != nil {
		z.log().Debug("Unable create request", zap.Error(err))
		return nil, err
	}
	for h, v := range z.headers {
		req.Header.Add(h, v)
	}
	return req, nil
}

func newPostRequest(ctx api_context.Context, url string, param interface{}, headers map[string]string) (req api_rpc.Request, err error) {
	// param
	p, err := json.Marshal(param)
	if err != nil {
		ctx.Log().Debug("Unable to marshal params", zap.Error(err))
		return nil, err
	}

	req = &requestImpl{
		ctx:     ctx,
		param:   p,
		url:     url,
		headers: headers,
		method:  "POST",
	}
	return req, nil
}
