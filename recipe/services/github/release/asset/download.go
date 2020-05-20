package asset

import (
	"github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release_asset"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Download struct {
	rc_recipe.RemarkExperimental
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
		l.Debug("download", esl.Any("asset", asset))
		ui.Progress(z.Downloading.With("File", asset.Name))
		path := filepath.Join(z.Path.Path(), asset.Name)
		if err := es_download.Download(c.Log(), asset.DownloadUrl, path); err != nil {
			return err
		}
	}
	return nil
}

func (z *Download) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("download", false)
	if err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(f)
	}()

	return rc_exec.ExecMock(c, &Download{}, func(r rc_recipe.Recipe) {
		m := r.(*Download)
		m.Owner = "watermint"
		m.Repository = "toolbox"
		m.Release = "65.4.233"
		m.Path = mo_path.NewFileSystemPath(f)
	})
}
