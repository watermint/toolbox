package es_context

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"net/http"
)

type Context interface {
	Name() string
	ClientHash() string
	Log() esl.Logger
	Capture() esl.Logger
}

type Classify func(res *http.Response) (success, failure es_response.Body)
