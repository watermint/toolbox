package release

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	rc_recipe.RemarkExperimental
	Owner      string
	Repository string
	Releases   rp_model.RowReport
	Peer       gh_conn.ConnGithubRepo
}

func (z *List) Preset() {
	z.Releases.SetModel(
		&mo_release.Release{},
		rp_model.HiddenColumns(
			"id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Releases.Open(); err != nil {
		return err
	}
	releases, err := sv_release.New(z.Peer.Context(), z.Owner, z.Repository).List()
	if err != nil {
		return err
	}
	for _, rel := range releases {
		z.Releases.Row(rel)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.Owner = "watermint"
		m.Repository = "toolbox"
	})
}
