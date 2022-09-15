package asset

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_release_asset"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release_asset"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	rc_recipe.RemarkExperimental
	Owner      string
	Repository string
	Release    string
	Peer       gh_conn.ConnGithubRepo
	Assets     rp_model.RowReport
}

func (z *List) Preset() {
	z.Assets.SetModel(
		&mo_release_asset.Asset{},
		rp_model.HiddenColumns("id"),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Assets.Open(); err != nil {
		return err
	}
	rel, err := sv_release.New(z.Peer.Client(), z.Owner, z.Repository).Get(z.Release)
	if err != nil {
		return err
	}

	assets, err := sv_release_asset.New(z.Peer.Client(), z.Owner, z.Repository, rel.Id).List()
	if err != nil {
		return err
	}

	for _, asset := range assets {
		z.Assets.Row(asset)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.Owner = "watermint"
		m.Repository = "toolbox"
		m.Release = "63.4.129"
	})
}
