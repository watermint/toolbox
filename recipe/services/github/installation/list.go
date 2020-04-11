package installation

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_installation"
	"github.com/watermint/toolbox/domain/github/service/sv_installation"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer          gh_conn.ConnGithubRepo
	Installations rp_model.RowReport
}

func (z *List) Preset() {
	z.Installations.SetModel(&mo_installation.Installation{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Installations.Open(); err != nil {
		return err
	}
	installations, err := sv_installation.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	for _, installation := range installations {
		z.Installations.Row(installation)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
