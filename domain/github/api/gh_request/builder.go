package gh_request

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"net/http"
)

func NewBuilder(ctl app_control.Control, token api_auth.Context) Builder {
	return Builder{
		ctl:   ctl,
		token: token,
	}
}

type Builder struct {
	ctl    app_control.Control
	token  api_auth.Context
	method string
	url    string
	data   api_request.RequestData
}

func (z Builder) Log() *zap.Logger {
	l := z.ctl.Log()
	if z.method != "" {
		l = l.With(zap.String("method", z.method))
	}
	if z.url != "" {
		l = l.With(zap.String("url", z.url))
	}
	if z.token != nil {
		l = l.With(zap.String("scope", z.token.Scope()))
	}
	return l
}

func (z Builder) Endpoint() string {
	return z.url
}

func (z Builder) Param() string {
	return string(z.data.ParamJson())
}

func (z Builder) ClientHash() string {
	var sr, st []string
	sr = []string{
		"m", z.method,
		"u", z.url,
	}
	if z.token != nil {
		st = []string{
			"p", z.token.PeerName(),
			"t", z.token.Token().AccessToken,
			"y", z.token.Scope(),
		}
	}

	return nw_client.ClientHash(sr, st)
}

func (z Builder) With(method, url string, data api_request.RequestData) Builder {
	z.method = method
	z.url = url
	z.data = data
	return z
}

func (z Builder) reqHeaders() map[string]string {
	headers := make(map[string]string)
	headers[api_request.ReqHeaderUserAgent] = app.UserAgent()
	if z.token != nil && !z.token.IsNoAuth() {
		headers[api_request.ReqHeaderAuthorization] = "token " + z.token.Token().AccessToken
	}

	// this will overwritten if a custom header provided thru request data
	headers[api_request.ReqHeaderContentType] = "application/json"
	for k, v := range z.data.Headers() {
		headers[k] = v
	}
	return headers
}

func (z Builder) reqContent() ut_io.ReadRewinder {
	if z.data.Content() != nil {
		return z.data.Content()
	}
	return ut_io.NewReadRewinderOnMemory(z.data.ParamJson())
}

func (z Builder) reqParamString() string {
	if z.data.Content() == nil {
		return ""
	}
	return z.data.ParamQuery()
}

func (z Builder) Build() (*http.Request, error) {
	l := z.Log()
	url := z.url + z.reqParamString()
	rc := z.reqContent()
	req, err := nw_client.NewHttpRequest(z.method, url, rc)
	if err != nil {
		l.Debug("Unable create request", zap.Error(err))
		return nil, err
	}
	headers := z.reqHeaders()
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.ContentLength = rc.Length()

	return req, nil
}
