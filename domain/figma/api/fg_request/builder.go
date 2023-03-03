package fg_request

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"net/http"
)

type Builder interface {
	api_request.Builder

	With(method, url string, data api_request.RequestData) Builder
}

type builderImpl struct {
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

func (z builderImpl) Build() (*http.Request, error) {
	//TODO implement me
	panic("implement me")
}

func (z builderImpl) WithData(data api_request.RequestDatum) api_request.Builder {
	//TODO implement me
	panic("implement me")
}

func (z builderImpl) Endpoint() string {
	//TODO implement me
	panic("implement me")
}

func (z builderImpl) Param() string {
	//TODO implement me
	panic("implement me")
}

func (z builderImpl) With(method, url string, data api_request.RequestData) Builder {
	//TODO implement me
	panic("implement me")
}
