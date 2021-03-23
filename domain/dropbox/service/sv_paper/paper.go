package sv_paper

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_paper"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Paper interface {
	Create(path mo_path.DropboxPath, format string, content []byte) (paper mo_paper.PaperUpdate, err error)
	Overwrite(path mo_path.DropboxPath, format string, content []byte) (paper mo_paper.PaperUpdate, err error)
	Append(path mo_path.DropboxPath, format string, content []byte) (paper mo_paper.PaperUpdate, err error)
	Prepend(path mo_path.DropboxPath, format string, content []byte) (paper mo_paper.PaperUpdate, err error)
}

func New(ctx dbx_context.Context) Paper {
	return &svPaper{
		ctx: ctx,
	}
}

type svPaper struct {
	ctx dbx_context.Context
}

func (z svPaper) Create(path mo_path.DropboxPath, format string, content []byte) (paper mo_paper.PaperUpdate, err error) {
	p := struct {
		Path         string `json:"path"`
		ImportFormat string `json:"import_format"`
	}{
		Path:         path.Path(),
		ImportFormat: format,
	}
	res := z.ctx.Post("files/paper/create", api_request.Param(p), api_request.Content(es_rewinder.NewReadRewinderOnMemory(content)))
	if err, f := res.Failure(); f {
		return mo_paper.PaperUpdate{}, err
	}
	if err := res.Success().Json().Model(&paper); err != nil {
		return mo_paper.PaperUpdate{}, err
	}
	return paper, nil
}

func (z svPaper) update(policy string, path mo_path.DropboxPath, format string, content []byte) (paper mo_paper.PaperUpdate, err error) {
	p := struct {
		Path            string `json:"path"`
		ImportFormat    string `json:"import_format"`
		DocUpdatePolicy string `json:"doc_update_policy"`
	}{
		Path:            path.Path(),
		ImportFormat:    format,
		DocUpdatePolicy: policy,
	}
	res := z.ctx.Post("files/paper/update", api_request.Param(p), api_request.Content(es_rewinder.NewReadRewinderOnMemory(content)))
	if err, f := res.Failure(); f {
		return mo_paper.PaperUpdate{}, err
	}
	if err := res.Success().Json().Model(&paper); err != nil {
		return mo_paper.PaperUpdate{}, err
	}
	return paper, nil
}

func (z svPaper) Overwrite(path mo_path.DropboxPath, format string, content []byte) (paper mo_paper.PaperUpdate, err error) {
	return z.update("overwrite", path, format, content)
}

func (z svPaper) Append(path mo_path.DropboxPath, format string, content []byte) (paper mo_paper.PaperUpdate, err error) {
	return z.update("append", path, format, content)
}

func (z svPaper) Prepend(path mo_path.DropboxPath, format string, content []byte) (paper mo_paper.PaperUpdate, err error) {
	return z.update("prepend", path, format, content)
}
