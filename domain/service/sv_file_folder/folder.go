package sv_file_folder

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
)

type Folder interface {
	// `files/create_folder`
	// options: auto_rename
	Create(path mo_path.Path) (entry mo_file.Entry, err error)
}

func New(ctx api_context.Context) Folder {
	return &folderImpl{
		ctx: ctx,
	}
}

type folderImpl struct {
	ctx api_context.Context
}

func (z *folderImpl) Create(path mo_path.Path) (entry mo_file.Entry, err error) {
	p := struct {
		Path       string `json:"path"`
		Autorename bool   `json:"autorename"`
	}{
		Path:       path.Path(),
		Autorename: false,
	}
	entry = &mo_file.Folder{}
	res, err := z.ctx.Rpc("files/create_folder_v2").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err := res.ModelWithPath(entry, "metadata"); err != nil {
		return nil, err
	}
	return entry, nil
}
