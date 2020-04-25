package sv_release_asset

import (
	"errors"
	"github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_release_asset"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
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
	res, err := z.ctx.Get(endpoint).Call()
	if err != nil {
		return nil, err
	}
	assets = make([]*mo_release_asset.Asset, 0)
	entries, found := res.Body().Json().Array()
	if !found {
		return nil, ErrorUnexpectedResponse
	}
	for _, entry := range entries {
		asset := &mo_release_asset.Asset{}
		if _, err := entry.Model(asset); err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

func (z *assetImpl) Upload(file mo_path.ExistingFileSystemPath) (asset *mo_release_asset.Asset, err error) {
	l := z.ctx.Log().With(
		zap.String("owner", z.owner),
		zap.String("repository", z.repository),
		zap.String("release", z.release),
		zap.String("path", file.Path()),
	)
	endpoint := "repos/" + z.owner + "/" + z.repository + "/releases/" + z.release + "/assets"
	p := struct {
		Name string `url:"name"`
	}{
		Name: filepath.Base(file.Path()),
	}
	contentType := mime.TypeByExtension(filepath.Ext(file.Path()))

	l.Debug("upload params",
		zap.String("endpoint", endpoint),
		zap.Any("param", p),
		zap.String("contentType", contentType))

	r, err := os.Open(file.Path())
	if err != nil {
		return nil, err
	}
	rr, err := ut_io.NewReadRewinder(r, 0)
	if err != nil {
		l.Debug("Unable to create read rewinder", zap.Error(err))
		return nil, err
	}
	defer r.Close()

	res, err := z.ctx.Upload(endpoint, rr).Param(p).
		Header(api_request.ReqHeaderContentType, contentType).Call()
	if err != nil {
		l.Debug("unable to upload", zap.Error(err))
		return nil, err
	}
	asset = &mo_release_asset.Asset{}
	if _, err := res.Body().Json().Model(asset); err != nil {
		l.Debug("failed to parse", zap.Error(err))
		return nil, err
	}
	return asset, nil
}
