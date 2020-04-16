package asset

import (
	"github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release_asset"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_download"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type Download struct {
	Path        mo_path.FileSystemPath
	Owner       string
	Repository  string
	Release     string
	Peer        gh_conn.ConnGithubRepo
	Downloads   rp_model.TransactionReport
	Downloading app_msg.Message
}

func (z *Download) Preset() {
	z.Downloads.SetModel(
		&UploadFile{},
		nil,
		rp_model.HiddenColumns("result.id"),
	)
}

func (z *Download) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()

	rel, err := sv_release.New(z.Peer.Context(), z.Owner, z.Repository).Get(z.Release)
	if err != nil {
		return err
	}

	sva := sv_release_asset.New(z.Peer.Context(), z.Owner, z.Repository, rel.Id)
	assets, err := sva.List()
	if err != nil {
		return err
	}

	err = os.MkdirAll(z.Path.Path(), 0755)
	if err != nil {
		return err
	}

	for _, asset := range assets {
		l.Debug("download", zap.Any("asset", asset))
		ui.Progress(z.Downloading.With("File", asset.Name))
		path := filepath.Join(z.Path.Path(), asset.Name)
		if err := ut_download.Download(c.Log(), asset.DownloadUrl, path); err != nil {
			return err
		}
	}
	return nil
}

func (z *Download) Test(c app_control.Control) error {
	return qt_errors.ErrorImplementMe
}
