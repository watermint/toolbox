package sv_file_folder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Folder interface {
	// `files/create_folder`
	// options: auto_rename
	Create(path mo_path.DropboxPath) (entry mo_file.Entry, err error)
}

func New(ctx dbx_context.Context) Folder {
	return &folderImpl{
		ctx: ctx,
	}
}

type folderImpl struct {
	ctx dbx_context.Context
}

func (z *folderImpl) Create(path mo_path.DropboxPath) (entry mo_file.Entry, err error) {
	p := struct {
		Path       string `json:"path"`
		Autorename bool   `json:"autorename"`
	}{
		Path:       path.Path(),
		Autorename: false,
	}
	res := z.ctx.Post("files/create_folder_v2", api_request.Param(p))
	if err, f := res.Failure(); f {
		return nil, err
	}
	entry = &mo_file.Folder{}
	err = res.Success().Json().FindModel("metadata", entry)
	return
}
