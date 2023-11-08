package nw_auth

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
	"github.com/watermint/toolbox/essentials/api/api_client"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"net/http"
)

type KeyMode string

const (
	KeyModeHeader KeyMode = "header"
	KeyModeParam  KeyMode = "param"
)

func NewKeyRestClient(entity api_auth.KeyEntity, repository api_auth.Repository, client nw_client.Rest, keyName string, keyMode KeyMode, keyFormatter func(string) string) nw_client.Rest {
	return &keyClient{
		entity:       entity,
		repository:   api_auth_repo.NewKey(repository),
		rest:         client,
		keyName:      keyName,
		keyMode:      keyMode,
		keyFormatter: keyFormatter,
	}
}

func NewKeyRequestBuilder(entity api_auth.KeyEntity, builder nw_client.RequestBuilder, keyName string, keyMode KeyMode, keyFormatter func(string) string) nw_client.RequestBuilder {
	return &keyRequestBuilder{
		builder:      builder,
		entity:       entity,
		keyName:      keyName,
		keyMode:      keyMode,
		keyFormatter: keyFormatter,
	}
}

type keyRequestBuilder struct {
	builder      nw_client.RequestBuilder
	entity       api_auth.KeyEntity
	keyName      string
	keyMode      KeyMode
	keyFormatter func(string) string
}

func (z keyRequestBuilder) Build() (*http.Request, error) {
	req, err := z.builder.Build()
	if err != nil {
		return nil, err
	}

	key := z.entity.Credential.Key
	if z.keyFormatter != nil {
		key = z.keyFormatter(key)
	}

	switch z.keyMode {
	case KeyModeHeader:
		req.Header.Add(z.keyName, key)
	case KeyModeParam:
		q := req.URL.Query()
		q.Add(z.keyName, key)
		req.URL.RawQuery = q.Encode()
	}
	return req, nil
}

func (z keyRequestBuilder) Endpoint() string {
	return z.builder.Endpoint()
}

func (z keyRequestBuilder) Param() string {
	return z.builder.Param()
}

type keyClient struct {
	entity       api_auth.KeyEntity
	repository   api_auth.KeyRepository
	rest         nw_client.Rest
	keyName      string
	keyMode      KeyMode
	keyFormatter func(string) string
}

func (z keyClient) Call(client api_client.Client, req nw_client.RequestBuilder) (res es_response.Response) {
	l := esl.Default().With(esl.String("endpoint", req.Endpoint()))
	abandon := func() {
		z.repository.Delete(z.entity.KeyName, z.entity.PeerName)
	}

	brq := NewKeyRequestBuilder(z.entity, req, z.keyName, z.keyMode, z.keyFormatter)
	res = z.rest.Call(client, brq)

	// abandon existing credential on auth error
	if res.Code() == 401 || res.IsAuthInvalidToken() {
		abandon()
		l.Debug("The response was invalid username/password")
		return es_response_impl.NewAuthErrorResponse(ErrorInvalidUsernameOrPassword, res)
	}

	return res
}
