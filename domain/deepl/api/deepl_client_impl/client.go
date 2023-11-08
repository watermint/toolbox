package deepl_client_impl

import (
	"github.com/watermint/toolbox/domain/deepl/api/deepl_client"
	"github.com/watermint/toolbox/domain/deepl/api/deepl_request"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_auth"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_rest_factory"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
	"strings"
)

var (
	V2Endpoint       = "https://api.deepl.com/v2/"
	V2EndpointFree   = "https://api-free.deepl.com/v2/"
	V2FreePeerPrefix = "free-"
)

func NewV2(name string, ctl app_control.Control, entity api_auth.KeyEntity) deepl_client.Client {
	client := nw_rest_factory.New(
		nw_rest_factory.Auth(func(client nw_client.Rest) (rest nw_client.Rest) {
			return nw_auth.NewKeyRestClient(entity, ctl.AuthRepository(), client, "Authorization", nw_auth.KeyModeHeader, func(key string) string {
				return "DeepL-Auth-Key " + key
			})
		}),
	)
	return &v2ClientImpl{
		peerName: name,
		client:   client,
		builder:  deepl_request.NewV2Builder(),
	}
}

type v2ClientImpl struct {
	peerName string
	client   nw_client.Rest
	ctl      app_control.Control
	builder  deepl_request.V2Builder
}

func (z v2ClientImpl) endpoint() string {
	if strings.HasPrefix(z.peerName, V2FreePeerPrefix) {
		return V2EndpointFree
	}
	return V2Endpoint
}

func (z v2ClientImpl) Name() string {
	return z.peerName
}

func (z v2ClientImpl) ClientHash() string {
	return z.builder.ClientHash()
}

func (z v2ClientImpl) Log() esl.Logger {
	return z.builder.Log()
}

func (z v2ClientImpl) Capture() esl.Logger {
	return z.ctl.Capture()
}

func (z v2ClientImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPost,
		z.endpoint()+endpoint,
		api_request.Combine(d),
	)
	return z.client.Call(&z, b)
}
