package issue

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_issue"
	"github.com/watermint/toolbox/domain/github/service/sv_issue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Owner      string
	Repository string
	Issues     rp_model.RowReport
	Peer       gh_conn.ConnGithubRepo
}

func (z *List) Preset() {
	z.Issues.SetModel(
		&mo_issue.Issue{},
		rp_model.HiddenColumns(
			"id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Issues.Open(); err != nil {
		return err
	}
	issues, err := sv_issue.New(z.Peer.Context(), z.Owner, z.Repository).List()
	if err != nil {
		return err
	}
	for _, issue := range issues {
		z.Issues.Row(issue)
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
