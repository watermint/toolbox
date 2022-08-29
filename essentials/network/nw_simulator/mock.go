package nw_simulator

import (
	"bytes"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/infra/api/api_client"
	"net/http"
)

type PanicClient struct {
}

func (p PanicClient) Call(ctx api_client.Client, req nw_client.RequestBuilder) (res es_response.Response) {
	panic("always panic!")
}

type MockApiContext struct {
}

func (z MockApiContext) Name() string {
	return ""
}

func (z MockApiContext) ClientHash() string {
	return ""
}

func (z MockApiContext) Log() esl.Logger {
	return esl.Default()
}

func (z MockApiContext) Capture() esl.Logger {
	return esl.Default()
}

type MockReqBuilder struct {
}

func (z MockReqBuilder) Build() (*http.Request, error) {
	return http.NewRequest("POST", z.Endpoint(), &bytes.Buffer{})
}

func (z MockReqBuilder) Endpoint() string {
	return "http://www.example.com"
}

func (z MockReqBuilder) Param() string {
	return ""
}
