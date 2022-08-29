package sv_file_restore

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Restore interface {
	Restore(path mo_path.DropboxPath, rev string) (entry mo_file.Entry, err error)
}

func New(ctx dbx_client.Client) Restore {
	return &restoreImpl{
		ctx: ctx,
	}
}

type restoreImpl struct {
	ctx dbx_client.Client
}

func (z *restoreImpl) Restore(path mo_path.DropboxPath, rev string) (entry mo_file.Entry, err error) {
	p := struct {
		Path string `json:"path"`
		Rev  string `json:"rev"`
	}{
		Path: path.Path(),
		Rev:  rev,
	}
	res := z.ctx.Post("files/restore", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	entry = &mo_file.Metadata{}
	err = res.Success().Json().Model(entry)
	return
}
