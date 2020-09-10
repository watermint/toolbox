package content

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_content"
	"github.com/watermint/toolbox/domain/github/service/sv_content"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Get struct {
	Owner      string
	Repository string
	Path       string
	Ref        mo_string.OptionalString
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
	opts := make([]sv_content.ContentOpt, 0)
	if z.Ref.IsExists() {
		opts = append(opts, sv_content.Ref(z.Ref.Value()))
	}
	cts, err := sv_content.New(z.Peer.Context(), z.Owner, z.Repository).Get(z.Path, opts...)
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
