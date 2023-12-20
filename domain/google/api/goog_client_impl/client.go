package goog_client_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response_impl"
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/domain/google/api/goog_request"
	"github.com/watermint/toolbox/domain/google/api/goog_response_impl"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_auth"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/essentials/network/nw_rest_factory"
	"github.com/watermint/toolbox/infra/control/app_apikey"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"net/http"
)

const (
	EndpointGoogleApis      EndpointType = "https://www.googleapis.com/"
	EndpointGoogleSheets    EndpointType = "https://sheets.googleapis.com/v4/"
	EndpointGoogleCalendar  EndpointType = "https://www.googleapis.com/calendar/v3/"
	EndpointGoogleTranslate EndpointType = "https://translate.googleapis.com/v3/"
)

type EndpointType string

func NewMock(endpoint EndpointType, name string, ctl app_control.Control) goog_client.Client {
	client := nw_rest_factory.New(
		nw_rest_factory.Mock())
	return &clientImpl{
		baseEndpoint: endpoint,
		name:         name,
		client:       client,
		ctl:          ctl,
		builder:      goog_request.NewBuilder(ctl, api_auth.NewNoAuthOAuthEntity()),
	}
}

func NewReplayMock(endpoint EndpointType, name string, ctl app_control.Control, rr []nw_replay.Response) goog_client.Client {
	client := nw_rest_factory.New(
		nw_rest_factory.Assert(dbx_response_impl.AssertResponse),
		nw_rest_factory.ReplayMock(rr))
	return &clientImpl{
		baseEndpoint: endpoint,
		name:         name,
		client:       client,
		ctl:          ctl,
		builder:      goog_request.NewBuilder(ctl, api_auth.NewNoAuthOAuthEntity()),
	}
}

func New(et EndpointType, name string, ctl app_control.Control, entity api_auth.OAuthEntity) goog_client.Client {
	var appData api_auth.OAuthAppData
	switch entity.KeyName {
	case app_definitions.ServiceGoogleCalendar:
		appData = goog_auth.Calendar
	case app_definitions.ServiceGoogleMail:
		appData = goog_auth.Mail
	case app_definitions.ServiceGoogleSheets:
		appData = goog_auth.Sheets
	case app_definitions.ServiceGoogleTranslate:
		appData = goog_auth.Translate
	default:
		panic("undefined app key type : " + entity.KeyName)
	}
	client := nw_rest_factory.New(
		nw_rest_factory.OAuthEntity(appData, func(appKey string) (clientId, clientSecret string) {
			return app_apikey.Resolve(ctl, appKey)
		}, entity),
		nw_rest_factory.Auth(func(client nw_client.Rest) (rest nw_client.Rest) {
			return nw_auth.NewOAuthRestClient(entity, ctl.AuthRepository(), client)
		}),
	)
	return &clientImpl{
		baseEndpoint: et,
		name:         name,
		client:       client,
		ctl:          ctl,
		builder:      goog_request.NewBuilder(ctl, entity),
	}
}

type clientImpl struct {
	baseEndpoint EndpointType
	name         string
	builder      goog_request.Builder
	client       nw_client.Rest
	ctl          app_control.Control
}

func (z clientImpl) Name() string {
	return z.name
}

func (z clientImpl) UI() app_ui.UI {
	return z.ctl.UI()
}

func (z clientImpl) Get(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodGet,
		string(z.baseEndpoint)+endpoint,
		api_request.Combine(d),
	)
	return goog_response_impl.New(z.client.Call(&z, b))
}

func (z clientImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPost,
		string(z.baseEndpoint)+endpoint,
		api_request.Combine(d),
	)
	return goog_response_impl.New(z.client.Call(&z, b))
}

func (z clientImpl) Put(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodPut,
		string(z.baseEndpoint)+endpoint,
		api_request.Combine(d),
	)
	return goog_response_impl.New(z.client.Call(&z, b))
}

func (z clientImpl) Delete(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	b := z.builder.With(
		http.MethodDelete,
		string(z.baseEndpoint)+endpoint,
		api_request.Combine(d),
	)
	return goog_response_impl.New(z.client.Call(&z, b))
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
