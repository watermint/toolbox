package deepl_request

import (
	"github.com/google/go-querystring/query"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"net/http"
)

func NewV2Builder() V2Builder {
	return &v2BuilderImpl{}
}

type V2Builder interface {
	api_request.Builder
	With(method, url string, data api_request.RequestData) V2Builder
}

type v2BuilderImpl struct {
	data   api_request.RequestData
	method string
	url    string
	entity api_auth.KeyEntity
}

func (z v2BuilderImpl) Log() esl.Logger {
	l := esl.Default()
	if z.method != "" {
		l = l.With(esl.String("method", z.method))
	}
	if z.url != "" {
		l = l.With(esl.String("url", z.url))
	}
	return l
}

func (z v2BuilderImpl) ClientHash() string {
	return nw_client.ClientHash(z.entity.HashSeed(), []string{
		"m", z.method,
		"u", z.url,
	})
}

func (z v2BuilderImpl) reqContent() es_rewinder.ReadRewinder {
	if z.data.Content() != nil {
		return z.data.Content()
	}
	return es_rewinder.NewReadRewinderOnMemory(z.data.ParamJson())
}

func (z v2BuilderImpl) Build() (*http.Request, error) {
	l := z.Log()
	rc := z.reqContent()
	qv, err := query.Values(z.data.Query())
	if err != nil {
		l.Debug("Unable to create query", esl.Error(err))
		return nil, err
	}
	reqUrl := z.url + "?" + qv.Encode()

	req, err := nw_client.NewHttpRequest(z.method, reqUrl, rc)
	if err != nil {
		l.Debug("Unable create request", esl.Error(err))
		return nil, err
	}

	headers := make(map[string]string)
	headers[api_request.ReqHeaderContentType] = "application/json"
	headers[api_request.ReqHeaderAccept] = "application/json"
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.ContentLength = rc.Length()

	return req, nil
}

func (z v2BuilderImpl) WithData(data api_request.RequestDatum) api_request.Builder {
	current := z.data.Data()
	allData := make([]api_request.RequestDatum, len(current)+1)
	copy(allData[:], current[:])
	allData[len(current)] = data

	z.data = api_request.Combine(allData)
	return z
}

func (z v2BuilderImpl) Endpoint() string {
	return z.url
}

func (z v2BuilderImpl) Param() string {
	return string(z.data.ParamJson())
}

func (z v2BuilderImpl) With(method, url string, data api_request.RequestData) V2Builder {
	z.method = method
	z.url = url
	z.data = data
	return z
}
