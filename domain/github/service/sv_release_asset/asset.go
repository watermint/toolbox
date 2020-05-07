package sv_release_asset

import (
	"errors"
	"github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_release_asset"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/api/api_request"
	"mime"
	"os"
	"path/filepath"
)

var (
	ErrorUnexpectedResponse = errors.New("unexpected response")
)

type Asset interface {
	List() (assets []*mo_release_asset.Asset, err error)
	Upload(file mo_path.ExistingFileSystemPath) (asset *mo_release_asset.Asset, err error)
}

func New(ctx gh_context.Context, owner, repository, release string) Asset {
	return &assetImpl{
		ctx:        ctx,
		owner:      owner,
		repository: repository,
		release:    release,
	}
}

type assetImpl struct {
	ctx        gh_context.Context
	owner      string
	repository string
	release    string
}

func (z *assetImpl) List() (assets []*mo_release_asset.Asset, err error) {
	endpoint := "repos/" + z.owner + "/" + z.repository + "/releases/" + z.release + "/assets"
	res := z.ctx.Get(endpoint)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	assets = make([]*mo_release_asset.Asset, 0)
	err = res.Success().Json().ArrayEach(func(e es_json.Json) error {
		asset := &mo_release_asset.Asset{}
		if err := e.Model(asset); err != nil {
			return err
		}
		assets = append(assets, asset)
		return nil
	})
	return
}

func (z *assetImpl) Upload(file mo_path.ExistingFileSystemPath) (asset *mo_release_asset.Asset, err error) {
	l := z.ctx.Log().With(
		es_log.String("owner", z.owner),
		es_log.String("repository", z.repository),
		es_log.String("release", z.release),
		es_log.String("path", file.Path()),
	)
	endpoint := "repos/" + z.owner + "/" + z.repository + "/releases/" + z.release + "/assets"
	p := struct {
		Name string `url:"name"`
	}{
		Name: filepath.Base(file.Path()),
	}
	contentType := mime.TypeByExtension(filepath.Ext(file.Path()))

	l.Debug("upload params",
		es_log.String("endpoint", endpoint),
		es_log.Any("param", p),
		es_log.String("contentType", contentType))

	r, err := os.Open(file.Path())
	if err != nil {
		return nil, err
	}
	rr, err := es_rewinder.NewReadRewinder(r, 0)
	if err != nil {
		l.Debug("Unable to create read rewinder", es_log.Error(err))
		return nil, err
	}
	defer r.Close()

	res := z.ctx.Upload(endpoint,
		api_request.Content(rr),
		api_request.Param(p),
		api_request.Header(api_request.ReqHeaderContentType, contentType))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	asset = &mo_release_asset.Asset{}
	err = res.Success().Json().Model(asset)
	return
}
