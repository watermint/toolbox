package nw_http

import (
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"net/http"
	"time"
)

type Mock struct {
}

func (z Mock) Call(clientHash string, endpoint string, req *http.Request) (res *http.Response, latency time.Duration, err error) {
	return nil, time.Nanosecond, qt_errors.ErrorMock
}
