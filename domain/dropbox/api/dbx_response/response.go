package dbx_response

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/http/es_response"
)

type Response interface {
	es_response.Response

	// Result of the API call.
	Result() es_json.Json

	// Dropbox API error. Returns empty instance if the response is success, or
	// the error type is not a dropbox api error.
	DropboxError() (err dbx_error.DropboxError)

	// The error about bad or expired token. This func will not return nil.
	ErrorAuth() dbx_error.ErrorAuth
	// The error about permission. This func will not return nil.
	ErrorAccess() dbx_error.ErrorAccess
	// The error about end point specific error. This func will not return nil.
	ErrorEndpointPath() dbx_error.ErrorEndpointPath
}
