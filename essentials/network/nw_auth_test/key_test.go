package nw_auth_test

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_auth"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_rest_factory"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"net/http"
	"net/url"
	"testing"
	"time"
)

const (
	testEndpoint = "https://httpbin.io/get"
)

type testApiBuilder struct {
	method string
}

func (z testApiBuilder) Build() (*http.Request, error) {
	return http.NewRequest(z.method, testEndpoint+"?otherParam="+time.Now().Format(time.RFC3339), nil)
}

func (z testApiBuilder) Endpoint() string {
	return testEndpoint
}

func (z testApiBuilder) Param() string {
	return ""
}

type testApiClient struct {
	client nw_client.Rest
	ctl    app_control.Control
}

func (z testApiClient) Name() string {
	return "test"
}

func (z testApiClient) ClientHash() string {
	return "test"
}

func (z testApiClient) Log() esl.Logger {
	return z.ctl.Log()
}

func (z testApiClient) Capture() esl.Logger {
	return z.ctl.Capture()
}

func (z testApiClient) Get(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	return es_response.NewProxy(z.client.Call(&z, &testApiBuilder{method: "GET"}))
}

func (z testApiClient) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	return es_response.NewProxy(z.client.Call(&z, &testApiBuilder{method: "POST"}))
}

func TestNewKeyRestClientWithHeader(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		header := "X-Test-Header"
		seed := time.Now().Format(time.RFC3339)
		entity := api_auth.KeyEntity{
			KeyName:  "testKeyName" + seed,
			PeerName: "testPeerName" + seed,
			Credential: api_auth.KeyCredential{
				Key: "testKey" + seed,
			},
			Description: "testDescription" + seed,
			Timestamp:   time.Now().Format(time.RFC3339),
		}
		authRepo := api_auth_repo.NewKey(ctl.AuthRepository())
		authRepo.Put(entity)
		client := nw_rest_factory.New(
			nw_rest_factory.Auth(func(client nw_client.Rest) (rest nw_client.Rest) {
				return nw_auth.NewKeyRestClient(entity, ctl.AuthRepository(), client, header, nw_auth.KeyModeHeader, func(s string) string {
					return "Bearer " + s
				})
			}),
		)
		api := testApiClient{
			client: client,
			ctl:    ctl,
		}
		res := api.Get("")
		if err, fail := res.Failure(); fail {
			t.Error(err)
		}
		h, found := res.Success().Json().FindString("headers." + header + ".0")
		if !found {
			t.Error("header not found")
		}
		if h != "Bearer "+entity.Credential.Key {
			t.Error("invalid key")
		}
	})
}

func TestNewKeyRestClientWithParam(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		paramName := "test_param"
		seed := time.Now().Format(time.RFC3339)
		entity := api_auth.KeyEntity{
			KeyName:  "testKeyName" + seed,
			PeerName: "testPeerName" + seed,
			Credential: api_auth.KeyCredential{
				Key: "testKey" + seed,
			},
			Description: "testDescription" + seed,
			Timestamp:   time.Now().Format(time.RFC3339),
		}
		authRepo := api_auth_repo.NewKey(ctl.AuthRepository())
		authRepo.Put(entity)
		client := nw_rest_factory.New(
			nw_rest_factory.Auth(func(client nw_client.Rest) (rest nw_client.Rest) {
				return nw_auth.NewKeyRestClient(entity, ctl.AuthRepository(), client, paramName, nw_auth.KeyModeParam, nil)
			}),
		)
		api := testApiClient{
			client: client,
			ctl:    ctl,
		}
		res := api.Get("")
		if err, fail := res.Failure(); fail {
			t.Error(err)
		}
		urlData, found := res.Success().Json().FindString("url")
		if !found {
			t.Error("param not found")
		}
		u, err := url.Parse(urlData)
		if err != nil {
			t.Error(err)
		}
		q := u.Query()
		if q.Get(paramName) != entity.Credential.Key {
			t.Error("invalid key")
		}
	})
}
