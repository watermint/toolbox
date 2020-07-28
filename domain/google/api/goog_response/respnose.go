package goog_response

import (
	"github.com/watermint/toolbox/domain/google/api/goog_error"
	"github.com/watermint/toolbox/essentials/http/es_response"
)

type Response interface {
	es_response.Response

	GoogleError() (err goog_error.GoogleError)
}
