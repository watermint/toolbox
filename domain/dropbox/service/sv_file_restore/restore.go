package sv_file_restore

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
)

type Restore interface {
	Restore(path mo_path.DropboxPath, rev string) (entry mo_file.Entry, err error)
}

func New(ctx dbx_context.Context) Restore {
	return &restoreImpl{
		ctx: ctx,
	}
}

type restoreImpl struct {
	ctx dbx_context.Context
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
	res, err := z.ctx.Post("files/restore").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if _, err := res.Body().Json().Model(entry); err != nil {
		return nil, err
	}
	return entry, nil
}
