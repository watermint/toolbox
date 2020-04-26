package dbx_request

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
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
	ctl        app_control.Control
	token      api_auth.Context
	asMemberId string
	asAdminId  string
	basePath   dbx_context.PathRoot
	method     string
	data       api_request.RequestData
	url        string
}

func (z Builder) Endpoint() string {
	return z.url
}

func (z Builder) NoAuth() Builder {
	z.token = nil
	return z
}
func (z Builder) AsMemberId(teamMemberId string) Builder {
	z.asMemberId = teamMemberId
	return z
}
func (z Builder) AsAdminId(teamMemberId string) Builder {
	z.asAdminId = teamMemberId
	return z
}
func (z Builder) WithPath(pathRoot dbx_context.PathRoot) Builder {
	z.basePath = pathRoot
	return z
}

func (z Builder) With(method, url string, data api_request.RequestData) Builder {
	z.method = method
	z.url = url
	z.data = data
	return z
}

func (z Builder) Log() *zap.Logger {
	l := z.ctl.Log()
	if z.asMemberId != "" {
		l = l.With(zap.String("asMemberId", z.asMemberId))
	}
	if z.asAdminId != "" {
		l = l.With(zap.String("asAdminId", z.asAdminId))
	}
	if z.basePath != nil {
		l = l.With(zap.Any("basePath", z.basePath))
	}
	return l
}

func (z Builder) ContentHash() string {
	var ss, sr, st, sp []string
	sr = []string{
		"m", z.method,
		"u", z.url,
	}
	ss = []string{
		"m", z.asMemberId,
		"a", z.asAdminId,
	}
	if z.token != nil {
		st = []string{
			"p", z.token.PeerName(),
			"t", z.token.Token().AccessToken,
			"y", z.token.Scope(),
		}
	}
	if z.basePath != nil {
		sp = []string{"p", z.basePath.Header()}
	}
	return nw_client.ClientHash(ss, sr, st, sp)
}

func (z Builder) Build() (*http.Request, error) {
	l := z.Log().With(zap.String("method", z.method), zap.String("url", z.url))
	rc := z.reqContent()
	req, err := nw_client.NewHttpRequest(z.method, z.url, rc)
	if err != nil {
		l.Debug("Unable to make request")
		return nil, err
	}
	headers := z.reqHeaders()
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.ContentLength = rc.Length()

	return req, nil
}

func (z Builder) reqHeaders() map[string]string {
	l := z.Log()

	headers := make(map[string]string)
	headers[api_request.ReqHeaderUserAgent] = app.UserAgent()
	if z.token != nil && !z.token.IsNoAuth() {
		headers[api_request.ReqHeaderAuthorization] = "Bearer " + z.token.Token().AccessToken
	}
	if z.asAdminId != "" {
		headers[api_request.ReqHeaderDropboxApiSelectAdmin] = z.asAdminId
	}
	if z.asMemberId != "" {
		headers[api_request.ReqHeaderDropboxApiSelectUser] = z.asMemberId
	}
	if z.basePath != nil {
		p, err := dbx_util.HeaderSafeJson(z.basePath)
		if err != nil {
			l.Debug("Unable to marshal base path", zap.Error(err))
		} else {
			headers[api_request.ReqHeaderDropboxApiPathRoot] = p
		}
	}
	if z.data.Content() != nil {
		p, err := dbx_util.HeaderSafeJson(z.data.Param())
		if err != nil {
			l.Debug("Unable to marshal params", zap.Error(err))
		} else {
			headers[api_request.ReqHeaderDropboxApiArg] = p
		}
		headers[api_request.ReqHeaderContentType] = "application/octet-stream"
	} else {
		headers[api_request.ReqHeaderContentType] = "application/json"
	}
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

func (z Builder) Param() string {
	return string(z.data.ParamJson())
}
