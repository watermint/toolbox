package sv_sharing

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/essentials/api/api_request"
)

type File interface {
	Resolve(file string) (meta *mo_file.Metadata, err error)
}

func New(ctx dbx_client.Client) File {
	return &fileImpl{
		ctx: ctx,
	}
}

type fileImpl struct {
	ctx dbx_client.Client
}

func (z fileImpl) Resolve(file string) (entry *mo_file.Metadata, err error) {
	p := struct {
		File string `json:"file"`
	}{
		File: file,
	}
	res := z.ctx.Post("sharing/get_file_metadata", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	entry = &mo_file.Metadata{}
	err = res.Success().Json().Model(entry)
	return
}
