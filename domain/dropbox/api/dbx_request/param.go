package dbx_request

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/infra/api/api_request"
)

func DropboxApiArg(p interface{}) (r api_request.RequestDatum, err error) {
	q, err := dbx_util.HeaderSafeJson(p)
	if err != nil {
		return nil, err
	}
	return api_request.Header(api_request.ReqHeaderDropboxApiArg, q), nil
}
