package dbx_response_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/http/es_response"
)

func New(res es_response.Response) dbx_response.Response {
	return &resImpl{
		Proxy:  es_response.NewProxy(res),
		result: nil,
	}
}

func NewAbort(res es_response.Response, err error) dbx_response.Response {
	return &resImpl{
		Proxy:  es_response.NewProxy(res),
		result: nil,
		abort:  err,
	}
}

type resImpl struct {
	es_response.Proxy

	result es_json.Json
	abort  error
}

func (z resImpl) DropboxError() (err dbx_error.ErrorInfo) {
	if z.IsSuccess() {
		return
	}
	_ = z.Alt().Json().Model(err)
	return
}

func (z resImpl) Failure() (error, bool) {
	if z.abort != nil {
		return z.abort, true
	}
	de := &dbx_error.ErrorInfo{}

	switch z.Code() {
	case 400: // bad input parameter
		return &dbx_error.ErrorBadRequest{Reason: z.Alt().BodyString()}, true

	case 401, 403, 409:
		if err := z.Alt().Json().Model(de); err == nil && de.ErrorSummary != "" {
			return de, true
		}
	}

	return z.Proxy.Failure()
}

func (z resImpl) Result() es_json.Json {
	if z.IsSuccess() {
		r := z.Header(dbx_context.DropboxApiResHeaderResult)
		if r != "" {
			return es_json.MustParseString(r)
		}
		return z.Success().Json()
	}
	return z.Alt().Json()
}

func (z resImpl) ErrorAuth() dbx_error.ErrorAuth {
	return dbx_error.NewErrorAuth(z.DropboxError())
}

func (z resImpl) ErrorAccess() dbx_error.ErrorAccess {
	return dbx_error.NewErrorAccess(z.DropboxError())
}

func (z resImpl) ErrorEndpointPath() dbx_error.ErrorEndpointPath {
	return dbx_error.NewErrorPath(z.DropboxError())
}
