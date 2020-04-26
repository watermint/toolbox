package nw_client

import (
	"crypto/sha256"
	"fmt"
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

type Rest interface {
	Call(ctx api_context.Context, req RequestBuilder) (res response.Response)
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

func NewGetRequest(url string, content ut_io.ReadRewinder) (*http.Request, error) {
	return NewHttpRequest(http.MethodGet, url, content)
}

func NewPostRequest(url string, content ut_io.ReadRewinder) (*http.Request, error) {
	return NewHttpRequest(http.MethodPost, url, content)
}

func NewHttpRequest(method, url string, content ut_io.ReadRewinder) (*http.Request, error) {
	l := app_root.Log()
	if err := content.Rewind(); err != nil {
		l.Debug("Unable to rewind", zap.Error(err))
		return nil, err
	}
	c := nw_bandwidth.WrapReader(content)
	req, err := http.NewRequest(method, url, c)
	if err != nil {
		l.Debug("unable to create request", zap.Error(err))
		return nil, err
	}
	return req, nil
}
