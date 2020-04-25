package sv_file_folder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
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
	entry = &mo_file.Folder{}
	res, err := z.ctx.Post("files/create_folder_v2").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if _, err := res.Success().Json().FindModel("metadata", entry); err != nil {
		return nil, err
	}
	return entry, nil
}
