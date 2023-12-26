package issue

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_issue"
	"github.com/watermint/toolbox/domain/github/service/sv_issue"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	rc_recipe.RemarkExperimental
	Owner      string
	Repository string
	Issues     rp_model.RowReport
	Peer       gh_conn.ConnGithubRepo
	State      mo_string.SelectString
	Labels     mo_string.OptionalString
	Since      mo_time.TimeOptional
	Filter     mo_string.SelectString
}

func (z *List) Preset() {
	z.Issues.SetModel(
		&mo_issue.Issue{},
		rp_model.HiddenColumns(
			"id",
		),
	)
	z.State.SetOptions("open", "open", "closed", "all")
	z.Filter.SetOptions("assigned", "assigned", "created", "mentioned", "subscribed", "repos", "all")
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Issues.Open(); err != nil {
		return err
	}
	opts := make([]sv_issue.ListOpt, 0)
	if z.Since.Ok() && !z.Since.IsZero() {
		opts = append(opts, sv_issue.ListSince(z.Since.Iso8601()))
	}
	if z.State.IsValid() {
		opts = append(opts, sv_issue.ListState(z.State.Value()))
	}
	if z.Labels.IsExists() {
		opts = append(opts, sv_issue.ListLabels(z.Labels.Value()))
	}
	if z.Filter.IsValid() {
		opts = append(opts, sv_issue.ListFilter(z.Filter.Value()))
	}
	issues, err := sv_issue.New(z.Peer.Client(), z.Owner, z.Repository).List(opts...)
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
