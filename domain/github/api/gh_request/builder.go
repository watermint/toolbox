package gh_request

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	api_request2 "github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
)

func NewBuilder(ctl app_control.Control, entity api_auth.OAuthEntity) Builder {
	return &builderImpl{
		ctl:    ctl,
		entity: entity,
	}
}

type Builder interface {
	api_request2.Builder
	With(method, url string, data api_request2.RequestData) Builder
}

type builderImpl struct {
	ctl    app_control.Control
	entity api_auth.OAuthEntity
	method string
	url    string
	data   api_request2.RequestData
}

func (z builderImpl) WithData(datum api_request2.RequestDatum) api_request2.Builder {
	return z.WithData(datum)
}

func (z builderImpl) Log() esl.Logger {
	l := z.ctl.Log()
	if z.method != "" {
		l = l.With(esl.String("method", z.method))
	}
	if z.url != "" {
		l = l.With(esl.String("url", z.url))
	}
	if !z.entity.IsNoAuth() {
		l = l.With(esl.Strings("scopes", z.entity.Scopes))
	}
	return l
}

func (z builderImpl) Endpoint() string {
	return z.url
}

func (z builderImpl) Param() string {
	return string(z.data.ParamJson())
}

func (z builderImpl) ClientHash() string {
	return nw_client.ClientHash(z.entity.HashSeed(), []string{
		"m", z.method,
		"u", z.url,
	})
}

func (z builderImpl) With(method, url string, data api_request2.RequestData) Builder {
	z.method = method
	z.url = url
	z.data = data
	return z
}

func (z builderImpl) reqHeaders() map[string]string {
	headers := make(map[string]string)
	headers[api_request2.ReqHeaderUserAgent] = app.UserAgent()
	if !z.entity.IsNoAuth() {
		headers[api_request2.ReqHeaderAuthorization] = "token " + z.entity.Token.AccessToken
	}

	// this will overwritten if a custom header provided thru request data
	headers[api_request2.ReqHeaderContentType] = "application/json"
	for k, v := range z.data.Headers() {
		headers[k] = v
	}
	return headers
}

func (z builderImpl) reqContent() es_rewinder.ReadRewinder {
	if z.data.Content() != nil {
		return z.data.Content()
	}
	return es_rewinder.NewReadRewinderOnMemory(z.data.ParamJson())
}

func (z builderImpl) reqParamString() string {
	//if z.data.Content() == nil {
	//	return ""
	//}
	return z.data.ParamQuery()
}

func (z builderImpl) Build() (*http.Request, error) {
	l := z.Log()
	url := z.url + z.reqParamString()
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
