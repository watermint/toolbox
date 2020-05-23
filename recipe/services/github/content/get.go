package content

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_content"
	"github.com/watermint/toolbox/domain/github/service/sv_content"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Get struct {
	Owner      string
	Repository string
	Path       string
	Content    rp_model.RowReport
	Peer       gh_conn.ConnGithubRepo
}

func (z *Get) Preset() {
	z.Content.SetModel(&mo_content.Content{})
}

func (z *Get) Exec(c app_control.Control) error {
	if err := z.Content.Open(); err != nil {
		return err
	}

	cts, err := sv_content.New(z.Peer.Context(), z.Owner, z.Repository).Get(z.Path)
	if err != nil {
		return err
	}

	if entries, found := cts.Dir(); found {
		for _, entry := range entries {
			z.Content.Row(entry)
		}
		return nil
	}
	if entry, found := cts.File(); found {
		z.Content.Row(entry)
		return nil
	}
	if entry, found := cts.Symlink(); found {
		z.Content.Row(entry)
		return nil
	}
	if entry, found := cts.Submodule(); found {
		z.Content.Row(entry)
		return nil
	}
	return nil
}

func (z *Get) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Get{}, func(r rc_recipe.Recipe) {
		m := r.(*Get)
		m.Owner = "watermint"
		m.Repository = "toolbox"
		m.Path = "/README.md"
	})
}
