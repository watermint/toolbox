package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/figma/api/fg_conn"
	"github.com/watermint/toolbox/domain/figma/model/mo_file"
	"github.com/watermint/toolbox/domain/figma/service/sv_project"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer      fg_conn.ConnFigmaApi
	Files     rp_model.RowReport
	ProjectId string
}

func (z *List) Preset() {
	z.Files.SetModel(&mo_file.File{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Files.Open(); err != nil {
		return err
	}
	if r, m := sv_project.VerifyProjectId(z.ProjectId); r != sv_project.VerifyProjectIdLooksOkay {
		c.UI().Error(m)
		return errors.New(c.UI().Text(m))
	}
	files, err := sv_project.New(z.Peer.Client()).Files(z.ProjectId)
	if err != nil {
		return err
	}
	for _, file := range files {
		z.Files.Row(file)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.ProjectId = "1234"
	})
}
