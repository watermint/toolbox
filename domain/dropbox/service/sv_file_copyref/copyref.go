package sv_file_copyref

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
)

var (
	ErrorUnexpectedFormat = errors.New("unexpected format")
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

	res, err := z.ctx.Post("files/copy_reference/get").Param(p).Call()
	if err != nil {
		return
	}
	ent := &mo_file.Metadata{}
	js := res.Body().Json()
	if _, err = js.FindModel("metadata", ent); err != nil {
		return
	}
	ref, found := js.FindString("copy_reference")
	if !found {
		return nil, "", "", ErrorUnexpectedFormat
	}
	expires, found = js.FindString("expires")
	if !found {
		return nil, "", "", ErrorUnexpectedFormat
	}
	return ent, ref, expires, nil
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
	res, err := z.ctx.Post("files/copy_reference/save").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if _, err = res.Body().Json().FindModel("metadata", entry); err != nil {
		return nil, err
	}
	return entry, nil
}
