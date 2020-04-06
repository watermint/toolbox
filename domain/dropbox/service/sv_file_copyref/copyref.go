package sv_file_copyref

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
)

type CopyRef interface {
	// `files/copy_reference/get`
	Resolve(path mo_path.DropboxPath) (entry mo_file.Entry, ref, expires string, err error)

	// `files/copy_reference/save`
	Save(path mo_path.DropboxPath, ref string) (entry mo_file.Entry, err error)
}

func New(ctx dbx_context.Context) CopyRef {
	return &copyRefImpl{
		ctx: ctx,
	}
}

type copyRefImpl struct {
	ctx dbx_context.Context
}

func (z *copyRefImpl) Resolve(path mo_path.DropboxPath) (entry mo_file.Entry, ref, expires string, err error) {
	p := struct {
		Path string `json:"path"`
	}{
		Path: path.Path(),
	}

	res, err := z.ctx.Rpc("files/copy_reference/get").Param(p).Call()
	if err != nil {
		return
	}
	js, err := res.Json()
	if err != nil {
		return
	}
	ent := &mo_file.Metadata{}
	if err = res.ModelWithPath(ent, "metadata"); err != nil {
		return
	}
	entry = ent
	ref = js.Get("copy_reference").String()
	expires = js.Get("expires").String()
	return
}

func (z *copyRefImpl) Save(path mo_path.DropboxPath, ref string) (entry mo_file.Entry, err error) {
	p := struct {
		CopyReference string `json:"copy_reference"`
		Path          string `json:"path"`
	}{
		CopyReference: ref,
		Path:          path.Path(),
	}

	entry = &mo_file.Metadata{}
	res, err := z.ctx.Rpc("files/copy_reference/save").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err = res.ModelWithPath(entry, "metadata"); err != nil {
		return nil, err
	}
	return entry, nil
}
