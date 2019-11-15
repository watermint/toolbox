package api_upload

import (
	"github.com/watermint/toolbox/infra/api/api_rpc"
	"io"
)

type Upload interface {
	Param(p interface{}) Upload
	Content(r io.Reader) Upload
	Call() (res api_rpc.Response, err error)
}
