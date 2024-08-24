package lifecycle

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/domain/github/model/mo_release_asset"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release_asset"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"time"
)

var (
	defaultAssetLifetimeDays = 31 * 6
)

type Assets struct {
	rc_recipe.RemarkSecret
	Peer    gh_conn.ConnGithubRepo
	Owner   string
	Repo    string
	Removed rp_model.RowReport
	Remove  bool
	Days    int
}

func (z *Assets) Preset() {
	z.Removed.SetModel(&mo_release_asset.Asset{})
	z.Owner = app_definitions.CoreRepositoryOwner
	z.Repo = app_definitions.CoreRepositoryName
	z.Days = defaultAssetLifetimeDays
}

func (z *Assets) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Removed.Open(); err != nil {
		return err
	}

	return sv_release.New(z.Peer.Client(), z.Owner, z.Repo).ListEach(func(release *mo_release.Release) error {
		assets, err := sv_release_asset.New(z.Peer.Client(), z.Owner, z.Repo, release.Id).List()
		if err != nil {
			l.Debug("Unable to list assets", esl.Error(err), esl.Any("release", release))
			return err
		}
		for _, asset := range assets {
			createdAt, err := time.Parse(time.RFC3339, asset.CreatedAt)
			if err != nil {
				l.Debug("Unable to parse created at", esl.Error(err), esl.Any("asset", asset))
				continue
			}
			if time.Since(createdAt).Hours() < float64(z.Days*24) {
				l.Debug("Skip asset", esl.Any("asset", asset))
				continue
			}

			if z.Remove {
				l.Info("Delete asset", esl.Any("asset", asset))
				if err := sv_release_asset.New(z.Peer.Client(), z.Owner, z.Repo, release.Id).Delete(asset.Id); err != nil {
					l.Warn("Unable to delete asset", esl.Error(err), esl.Any("asset", asset))
					return err
				}
			}

			z.Removed.Row(asset)
		}
		return nil
	})
}

func (z *Assets) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Assets{}, func(r rc_recipe.Recipe) {
		m := r.(*Assets)
		m.Owner = "watermint"
		m.Repo = "toolbox"
	})
}
