package dbx_response_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/http/response"
)

func New(res response.Response) dbx_response.Response {
	return &resImpl{
		Proxy:  response.NewProxy(res),
		result: nil,
	}
}

func NewAbort(res response.Response, err error) dbx_response.Response {
	return &resImpl{
		Proxy:  response.NewProxy(res),
		result: nil,
		abort:  err,
	}
}

type resImpl struct {
	response.Proxy

	result tjson.Json
	abort  error
}

func (z resImpl) DropboxError() (err dbx_error.DropboxError) {
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
	de := &dbx_error.DropboxError{}
	if err := z.Alt().Json().Model(de); err == nil && de.ErrorSummary != "" {
		return de, true
	}
	return z.Proxy.Failure()
}

func (z resImpl) Result() tjson.Json {
	if z.IsSuccess() {
		r := z.Header(dbx_context.DropboxApiResHeaderResult)
		if r != "" {
			return tjson.MustParseString(r)
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
