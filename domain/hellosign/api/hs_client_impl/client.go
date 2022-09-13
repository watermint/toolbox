package hs_client_impl

import (
	"github.com/watermint/toolbox/domain/hellosign/api/hs_client"
	"github.com/watermint/toolbox/domain/hellosign/api/hs_request"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_auth"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_rest_factory"
	"github.com/watermint/toolbox/infra/control/app_control"
	"net/http"
)

const (
	Endpoint = "https://api.hellosign.com/v3/"
)

func NewMock(name string, ctl app_control.Control) hs_client.Client {
	return &clientImpl{
		peerName: name,
		client:   nw_rest_factory.New(nw_rest_factory.Mock()),
		ctl:      ctl,
	}
}

func New(name string, ctl app_control.Control, entity api_auth.BasicEntity) hs_client.Client {
	client := nw_rest_factory.New(
		nw_rest_factory.Auth(func(client nw_client.Rest) (rest nw_client.Rest) {
			return nw_auth.NewBasicRestClient(entity, ctl.AuthRepository(), client)
		}),
	)
	return &clientImpl{
		peerName: name,
		client:   client,
		ctl:      ctl,
		builder:  hs_request.NewBuilder(),
	}
}

type clientImpl struct {
	peerName string
	client   nw_client.Rest
	ctl      app_control.Control
	builder  hs_request.Builder
}

func (z clientImpl) Name() string {
	return z.peerName
}

func (z clientImpl) ClientHash() string {
	return z.builder.ClientHash()
}

func (z clientImpl) Log() esl.Logger {
	return z.builder.Log()
}

func (z clientImpl) Capture() esl.Logger {
	return z.ctl.Capture()
}

func (z clientImpl) Get(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodGet,
		Endpoint+endpoint,
		api_request.Combine(d),
	)
	return z.client.Call(&z, b)
}

func (z clientImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPost,
		Endpoint+endpoint,
		api_request.Combine(d),
	)
	return z.client.Call(&z, b)
}

func (z clientImpl) Put(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPut,
		Endpoint+endpoint,
		api_request.Combine(d),
	)
	return z.client.Call(&z, b)
}
