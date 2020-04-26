package sv_file_url

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	url2 "net/url"
	"path/filepath"
)

type Url interface {
	Save(path mo_path.DropboxPath, url string) (entry mo_file.Entry, err error)
}

func New(ctx dbx_context.Context) Url {
	return &urlImpl{
		ctx: ctx,
	}
}

// Path with filename that parsed from the URL.
func PathWithName(base mo_path.DropboxPath, url string) (path mo_path.DropboxPath) {
	u, err := url2.Parse(url)
	if err != nil {
		app_root.Log().Debug("Unable to parse url", zap.Error(err), zap.String("url", url))
		n := filepath.Base(url)
		return base.ChildPath(n)
	}
	if u.Path == "" {
		return base
	}
	return base.ChildPath(filepath.Base(u.Path))
}

type urlImpl struct {
	ctx dbx_context.Context
}

func (z *urlImpl) Save(path mo_path.DropboxPath, url string) (entry mo_file.Entry, err error) {
	p := struct {
		Path string `json:"path"`
		Url  string `json:"url"`
	}{
		Path: path.Path(),
		Url:  url,
	}

	meta := &mo_file.Metadata{}
	res := z.ctx.Async("files/save_url", api_request.Param(p)).Call(
		dbx_async.Status("files/save_url/check_job_status"))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	err = res.Success().Json().Model(meta)
	meta.EntryTag = "file" // overwrite 'complete' tag
	entry = meta
	return meta, err
}
