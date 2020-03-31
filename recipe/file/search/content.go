package search

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Content struct {
	Peer      rc_conn.ConnUserFile
	Path      string
	Query     string
	Extension string
	Category  string
	Matches   rp_model.RowReport
}

func (z *Content) Exec(c app_control.Control) error {
	so := make([]sv_file.SearchOpt, 0)
	so = append(so, sv_file.SearchIncludeHighlights())
	if z.Extension != "" {
		so = append(so, sv_file.SearchFileExtension(z.Extension))
	}
	if z.Category != "" {
		so = append(so, sv_file.SearchCategories(z.Category))
	}
	if z.Path != "" {
		so = append(so, sv_file.SearchPath(mo_path.NewDropboxPath(z.Path)))
	}

	if err := z.Matches.Open(); err != nil {
		return err
	}

	matches, err := sv_file.NewFiles(z.Peer.Context()).Search(z.Query, so...)
	if err != nil {
		return err
	}

	for _, m := range matches {
		z.Matches.Row(m.Highlighted())
	}
	return nil
}

func (z *Content) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Content{}, func(r rc_recipe.Recipe) {
		m := r.(*Content)
		m.Query = "watermint"
	})
}

func (z *Content) Preset() {
	z.Matches.SetModel(&mo_file.MatchHighlighted{}, rp_model.HiddenColumns("name", "path_lower"))
}
