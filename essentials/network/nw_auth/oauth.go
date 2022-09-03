package nw_auth

import (
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/network/nw_client"
)

func NewOAuthRestClient(entity api_auth2.OAuthEntity, repository api_auth2.Repository, client nw_client.Rest) nw_client.Rest {
	return &oauthClient{
		entity:     entity,
		repository: api_auth_repo.NewOAuth(repository),
		rest:       client,
	}
}

type oauthClient struct {
	entity     api_auth2.OAuthEntity
	repository api_auth2.OAuthRepository
	rest       nw_client.Rest
}

func (z oauthClient) Call(client api_client.Client, req nw_client.RequestBuilder) (res es_response.Response) {
	res = z.rest.Call(client, req)
	if res.IsSuccess() {
		return res
	}

	// abandon existing token on auth error
	if res.Code() == 401 || res.IsAuthInvalidToken() {
		z.repository.Delete(z.entity.KeyName, z.entity.Scopes, z.entity.PeerName)
	}
	return res
}
