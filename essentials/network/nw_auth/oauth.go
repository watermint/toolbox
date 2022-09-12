package nw_auth

import (
	"errors"
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"strings"
)

var (
	ErrorInvalidOrExpiredRefreshToken = errors.New("invalid or expired refresh token")
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

	l := esl.Default().With(esl.String("reqEndpoint", req.Endpoint()))
	abandon := func() {
		z.repository.Delete(z.entity.KeyName, z.entity.Scopes, z.entity.PeerName)
	}

	if err, fail := res.Failure(); fail {
		errText := err.Error()
		if strings.Contains(errText, "oauth2: cannot fetch token:") {
			abandon()
			l.Debug("The error response contains oauth2 error", esl.Error(err))
			return es_response_impl.NewAuthErrorResponse(ErrorInvalidOrExpiredRefreshToken, res)
		}
	}
	// abandon existing token on auth error
	if res.Code() == 401 || res.IsAuthInvalidToken() {
		abandon()
		l.Debug("The response was invalid token")
	}
	return res
}
