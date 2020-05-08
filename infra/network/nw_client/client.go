package nw_client

import (
	"crypto/sha256"
	"fmt"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"net/http"
	"strings"
	"time"
)

type Rest interface {
	Call(ctx api_context.Context, req RequestBuilder) (res es_response.Response)
}

type Http interface {
	Call(clientHash string, endpoint string, req *http.Request) (res *http.Response, latency time.Duration, err error)
}

type RequestBuilder interface {
	// Create new http request
	Build() (*http.Request, error)

	// Identifier of endpoint. That could be url or part of url.
	// This will be used for QoS control.
	Endpoint() string

	// String form of parameters. This will be used for logging.
	Param() string
}

func ClientHash(seeds ...[]string) string {
	all := make([]string, 0)
	for _, s := range seeds {
		all = append(all, s...)
	}

	return fmt.Sprintf("%x", sha256.Sum224([]byte(strings.Join(all, ","))))
}

func NewGetRequest(url string, content es_rewinder.ReadRewinder) (*http.Request, error) {
	return NewHttpRequest(http.MethodGet, url, content)
}

func NewPostRequest(url string, content es_rewinder.ReadRewinder) (*http.Request, error) {
	return NewHttpRequest(http.MethodPost, url, content)
}

func NewHttpRequest(method, url string, content es_rewinder.ReadRewinder) (*http.Request, error) {
	l := esl.Default()
	if err := content.Rewind(); err != nil {
		l.Debug("Unable to rewind", esl.Error(err))
		return nil, err
	}
	c := nw_bandwidth.WrapReader(content)
	req, err := http.NewRequest(method, url, c)
	if err != nil {
		l.Debug("unable to create request", esl.Error(err))
		return nil, err
	}
	return req, nil
}
