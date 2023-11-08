package nw_auth

import (
	"errors"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"net/http"
)

const (
	DefaultHeaderKey = "Authorization"
)

var (
	ErrorInvalidUsernameOrPassword = errors.New("invalid username or password")
)

func NewBasicRestClient(entity api_auth.BasicEntity, repository api_auth.Repository, client nw_client.Rest) nw_client.Rest {
	return &basicClient{
		entity:     entity,
		repository: api_auth_repo.NewBasic(repository),
		rest:       client,
		headerKey:  DefaultHeaderKey,
	}
}

func NewBasicRestClientWithHeaderKey(entity api_auth.BasicEntity, repository api_auth.Repository, client nw_client.Rest, headerKey string) nw_client.Rest {
	return &basicClient{
		entity:     entity,
		repository: api_auth_repo.NewBasic(repository),
		rest:       client,
		headerKey:  headerKey,
	}
}

func NewBasicRequestBuilder(entity api_auth.BasicEntity, builder nw_client.RequestBuilder, headerKey string) nw_client.RequestBuilder {
	return &basicRequestBuilder{
		builder:   builder,
		entity:    entity,
		headerKey: headerKey,
	}
}

type basicRequestBuilder struct {
	builder   nw_client.RequestBuilder
	entity    api_auth.BasicEntity
	headerKey string
}

func (z basicRequestBuilder) Build() (*http.Request, error) {
	req, err := z.builder.Build()
	if err != nil {
		return nil, err
	}
	req.Header.Add(z.headerKey, z.entity.Credential.HeaderValue())
	return req, nil
}

func (z basicRequestBuilder) Endpoint() string {
	return z.builder.Endpoint()
}

func (z basicRequestBuilder) Param() string {
	return z.builder.Param()
}

type basicClient struct {
	entity     api_auth.BasicEntity
	repository api_auth.BasicRepository
	rest       nw_client.Rest
	headerKey  string
}

func (z basicClient) Call(client api_client.Client, req nw_client.RequestBuilder) (res es_response.Response) {
	l := esl.Default().With(esl.String("endpoint", req.Endpoint()))
	abandon := func() {
		z.repository.Delete(z.entity.KeyName, z.entity.PeerName)
	}

	brq := NewBasicRequestBuilder(z.entity, req, z.headerKey)
	res = z.rest.Call(client, brq)

	// abandon existing credential on auth error
	if res.Code() == 401 || res.IsAuthInvalidToken() {
		abandon()
		l.Debug("The response was invalid username/password")
		return es_response_impl.NewAuthErrorResponse(ErrorInvalidUsernameOrPassword, res)
	}

	return res
}
