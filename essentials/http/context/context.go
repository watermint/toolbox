package context

import (
	"github.com/watermint/toolbox/essentials/http/response"
	"go.uber.org/zap"
	"net/http"
)

type Context interface {
	ClientHash() string
	Log() *zap.Logger
	Capture() *zap.Logger
}

type Classify func(res *http.Response) (success, failure response.Body)
