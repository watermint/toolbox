package fg_request

import (
	"github.com/google/go-querystring/query"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
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
		entity: api_auth.BasicEntity{},
	}
}

type Builder interface {
	api_request.Builder

	With(method, url string, data api_request.RequestData) Builder
}

type builderImpl struct {
	ctl    app_control.Control
	data   api_request.RequestData
	entity api_auth.BasicEntity
	method string
	url    string
}

func (z builderImpl) Log() esl.Logger {
	l := esl.Default()
	if z.method != "" {
		l = l.With(esl.String("method", z.method))
	}
	if z.url != "" {
		l = l.With(esl.String("url", z.url))
	}
	return l
}

func (z builderImpl) ClientHash() string {
	return nw_client.ClientHash(z.entity.HashSeed(), []string{
		"m", z.method,
		"u", z.url,
	})
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

	for k, v := range z.data.Headers() {
		headers[k] = v
	}
	return
}

func (z builderImpl) Build() (*http.Request, error) {
	l := z.ctl.Log()
	qv, err := query.Values(z.data.Query())
	if err != nil {
		l.Debug("Unable to create query", esl.Error(err))
		return nil, err
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

func (z builderImpl) WithData(data api_request.RequestDatum) api_request.Builder {
	current := z.data.Data()
	allData := make([]api_request.RequestDatum, len(current)+1)
	copy(allData[:], current[:])
	allData[len(current)] = data

	z.data = api_request.Combine(allData)
	return z
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
