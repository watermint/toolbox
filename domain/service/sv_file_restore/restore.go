package sv_file_restore

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
)

type Restore interface {
	Restore(path mo_path.DropboxPath, rev string) (entry mo_file.Entry, err error)
}

func New(ctx api_context.Context) Restore {
	return &restoreImpl{
		ctx: ctx,
	}
}

type restoreImpl struct {
	ctx api_context.Context
}

func (z *restoreImpl) Restore(path mo_path.DropboxPath, rev string) (entry mo_file.Entry, err error) {
	p := struct {
		Path string `json:"path"`
		Rev  string `json:"rev"`
	}{
		Path: path.Path(),
		Rev:  rev,
	}
	entry = &mo_file.Metadata{}
	res, err := z.ctx.Rpc("files/restore").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err := res.Model(entry); err != nil {
		return nil, err
	}
	return entry, nil
}
