package project

import (
	"errors"
	"github.com/watermint/toolbox/domain/figma/api/fg_conn"
	"github.com/watermint/toolbox/domain/figma/model/mo_project"
	"github.com/watermint/toolbox/domain/figma/service/sv_project"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer     fg_conn.ConnFigmaApi
	Projects rp_model.RowReport
	TeamId   string
}

func (z *List) Preset() {
	z.Projects.SetModel(&mo_project.Project{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Projects.Open(); err != nil {
		return err
	}
	if r, m := sv_project.VerifyTeamId(z.TeamId); r != sv_project.VerifyTeamIdLooksOkay {
		c.UI().Error(m)
		return errors.New(c.UI().Text(m))
	}

	projects, err := sv_project.New(z.Peer.Client()).List(z.TeamId)
	if err != nil {
		return err
	}
	for _, project := range projects {
		z.Projects.Row(project)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.TeamId = "1234"
	})
}
