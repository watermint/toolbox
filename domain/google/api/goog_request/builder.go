package goog_request

import (
	"github.com/google/go-querystring/query"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"net/http"
	url2 "net/url"
)

func NewBuilder(ctl app_control.Control, entity api_auth.OAuthEntity) Builder {
	return &builderImpl{
		disablePretty: true,
		ctl:           ctl,
		entity:        entity,
	}
}

type Builder interface {
	api_request.Builder
	With(method, url string, data api_request.RequestData) Builder
}

type builderImpl struct {
	disablePretty bool
	ctl           app_control.Control
	entity        api_auth.OAuthEntity
	method        string
	url           string
	data          api_request.RequestData
}

func (z builderImpl) WithData(data api_request.RequestDatum) api_request.Builder {
	z.data = z.data.WithDatum(data)
	return z
}

func (z builderImpl) With(method, url string, data api_request.RequestData) Builder {
	z.method = method
	z.url = url
	z.data = data
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
	if !z.entity.IsNoAuth() {
		l = l.With(esl.Strings("scopes", z.entity.Scopes))
	}
	return l
}

func (z builderImpl) ClientHash() string {
	return nw_client.ClientHash(z.entity.HashSeed(), []string{
		"m", z.method,
		"u", z.url,
	})
}

func (z builderImpl) Endpoint() string {
	return z.url
}

func (z builderImpl) Param() string {
	return string(z.data.ParamJson())
}

func (z builderImpl) reqHeaders() map[string]string {
	headers := make(map[string]string)
	headers[api_request.ReqHeaderUserAgent] = app_definitions.UserAgent()
	if !z.entity.IsNoAuth() {
		headers[api_request.ReqHeaderAuthorization] = "token " + z.entity.Token.AccessToken
	}

	// this will overwritten if a custom header provided thru request data
	headers[api_request.ReqHeaderContentType] = "application/json"
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

func (z builderImpl) Build() (req *http.Request, err error) {
	l := z.Log()
	var url string
	if z.disablePretty {
		var queryValue url2.Values
		if z.data.Query() == nil {
			queryValue, err = query.Values(struct {
				Pretty bool `url:"$prettyPrint"`
			}{
				Pretty: false,
			})
			if err != nil {
				l.Debug("Unable to create query params", esl.Error(err))
				return nil, err
			}
		} else {
			queryValue, err = query.Values(z.data.Query())
			if err != nil {
				l.Debug("Unable to create query params", esl.Error(err))
				return nil, err
			}
			queryValue.Add("$prettyPrint", "false")
		}
		url = z.url + "?" + queryValue.Encode()
	} else {
		url = z.url + z.data.ParamQuery()
	}
	rc := z.reqContent()
	req, err = nw_client.NewHttpRequest(z.method, url, rc)
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
