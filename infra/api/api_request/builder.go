package api_request

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"net/http"
)

type Builder interface {
	Log() esl.Logger
	ClientHash() string
	Build() (*http.Request, error)
	Endpoint() string
	Param() string
}
