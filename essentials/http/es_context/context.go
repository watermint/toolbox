package es_context

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"net/http"
)

type Context interface {
	ClientHash() string
	Log() es_log.Logger
	Capture() es_log.Logger
}

type Classify func(res *http.Response) (success, failure es_response.Body)
