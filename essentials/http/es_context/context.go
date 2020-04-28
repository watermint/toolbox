package es_context

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"go.uber.org/zap"
	"net/http"
)

type Context interface {
	ClientHash() string
	Log() *zap.Logger
	Capture() *zap.Logger
}

type Classify func(res *http.Response) (success, failure es_response.Body)
