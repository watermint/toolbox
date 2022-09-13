package nw_auth

import (
	"errors"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
)

var (
	ErrorInvalidUsernameOrPassword = errors.New("invalid username or password")
)

func NewBasicRestClient(entity api_auth.BasicEntity, repository api_auth.Repository, client nw_client.Rest) nw_client.Rest {
	return &basicClient{
		entity:     entity,
		repository: api_auth_repo.NewBasic(repository),
		rest:       client,
	}
}

type basicClient struct {
	entity     api_auth.BasicEntity
	repository api_auth.BasicRepository
	rest       nw_client.Rest
}

func (z basicClient) Call(client api_client.Client, req nw_client.RequestBuilder) (res es_response.Response) {
	l := esl.Default().With(esl.String("endpoint", req.Endpoint()))
	mrq1, ok := req.(nw_client.MutableRequestBuilder)
	if !ok {
		panic("Basic Auth client requires MutableRequestBuilder implementation")
	}
	abandon := func() {
		z.repository.Delete(z.entity.KeyName, z.entity.PeerName)
	}

	mrq2 := mrq1.WithData(api_request.Header("Authorization", z.entity.Credential.HeaderValue()))
	res = z.rest.Call(client, mrq2)

	// abandon existing credential on auth error
	if res.Code() == 401 || res.IsAuthInvalidToken() {
		abandon()
		l.Debug("The response was invalid username/password")
		return es_response_impl.NewAuthErrorResponse(ErrorInvalidUsernameOrPassword, res)
	}

	return res
}
