package sv_file_url

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
)

type Url interface {
	Save(path mo_path.Path, url string) (entry mo_file.Entry, err error)
}

func New(ctx api_context.Context) Url {
	return &urlImpl{
		ctx: ctx,
	}
}

type urlImpl struct {
	ctx api_context.Context
}

func (z *urlImpl) Save(path mo_path.Path, url string) (entry mo_file.Entry, err error) {
	p := struct {
		Path string `json:"path"`
		Url  string `json:"url"`
	}{
		Path: path.Path(),
		Url:  url,
	}

	meta := &mo_file.Metadata{}
	entry = meta
	res, err := z.ctx.Async("files/save_url").
		Status("files/save_url/check_job_status").
		Param(p).
		Call()
	if err != nil {
		return nil, err
	}
	if err = res.Model(entry); err != nil {
		return nil, err
	}
	meta.EntryTag = "file" // overwrite 'complete' tag
	return entry, nil
}
