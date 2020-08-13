package as_request

import (
	"github.com/google/go-querystring/query"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
	"strconv"
	"strings"
)

func NewBuilder(ctl app_control.Control, token api_auth.Context) Builder {
	return &builderImpl{
		ctl:   ctl,
		token: token,
	}
}

type Builder interface {
	api_request.Builder
	With(method, url string, data api_request.RequestData) Builder
	WithOffset(limit int, offset string) Builder
}

type builderImpl struct {
	ctl    app_control.Control
	token  api_auth.Context
	method string
	url    string
	limit  int
	offset string
	data   api_request.RequestData
}

func (z builderImpl) WithOffset(limit int, offset string) Builder {
	z.limit = limit
	z.offset = offset
	return z
}

func (z builderImpl) Log() esl.Logger {
	l := z.ctl.Log()
	if z.method != "" {
		l = l.With(esl.String("method", z.method))
	}
	if z.url != "" {
		l = l.With(esl.String("url", z.url))
	}
	if z.token != nil {
		l = l.With(esl.Strings("scopes", z.token.Scopes()))
	}
	return l
}

func (z builderImpl) ClientHash() string {
	var sr, st []string
	sr = []string{
		"m", z.method,
		"u", z.url,
	}
	if z.token != nil {
		st = []string{
			"p", z.token.PeerName(),
			"t", z.token.Token().AccessToken,
			"y", strings.Join(z.token.Scopes(), ","),
		}
	}

	return nw_client.ClientHash(sr, st)
}

func (z builderImpl) reqContent() es_rewinder.ReadRewinder {
	if z.data.Content() != nil {
		return z.data.Content()
	}
	return es_rewinder.NewReadRewinderOnMemory(z.data.ParamJson())
}

func (z builderImpl) reqHeaders() (headers map[string]string) {
	headers = make(map[string]string)
	headers[api_request.ReqHeaderUserAgent] = app.UserAgent()
	headers[api_request.ReqHeaderContentType] = "application/json"
	headers[api_request.ReqHeaderAccept] = "application/json"
	if z.token != nil {
		headers[api_request.ReqHeaderAuthorization] = "Bearer " + z.token.Token().AccessToken
	} else {
		headers[api_request.ReqHeaderAuthorization] = "MOCK_CALL"
	}

	for k, v := range z.data.Headers() {
		headers[k] = v
	}
	return
}

func (z builderImpl) Build() (*http.Request, error) {
	l := z.Log()
	qv, err := query.Values(z.data.Query())
	if err != nil {
		l.Debug("Unable to create query", esl.Error(err))
		return nil, err
	}
	qv.Add("opt_pretty", "false")
	if 0 < z.limit {
		qv.Add("limit", strconv.FormatInt(int64(z.limit), 10))
		if z.offset != "" {
			qv.Add("offset", z.offset)
		}
	}
	url := z.url + "?" + qv.Encode()

	rc := z.reqContent()
	req, err := nw_client.NewHttpRequest(z.method, url, rc)
	if err != nil {
		l.Debug("Unable create request", esl.Error(err))
		return nil, err
	}
	headers := z.reqHeaders()
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.ContentLength = rc.Length()

	return req, nil
}

func (z builderImpl) Endpoint() string {
	return z.url
}

func (z builderImpl) Param() string {
	return string(z.data.ParamJson())
}

func (z builderImpl) With(method, url string, data api_request.RequestData) Builder {
	z.method = method
	z.url = url
	z.data = data
	return z
}
