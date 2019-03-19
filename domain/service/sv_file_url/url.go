package sv_file_url

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
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

	entry = &mo_file.Metadata{}
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
	return entry, nil
}
