package search

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Name struct {
	Peer      rc_conn.ConnUserFile
	Path      string
	Query     string
	Extension string
	Category  string
	Matches   rp_model.RowReport
}

func (z *Name) Exec(c app_control.Control) error {
	so := make([]sv_file.SearchOpt, 0)
	so = append(so, sv_file.SearchFileNameOnly())
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

func (z *Name) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Name{}, func(r rc_recipe.Recipe) {
		m := r.(*Name)
		m.Query = "watermint"
	})
}

func (z *Name) Preset() {
	z.Matches.SetModel(&mo_file.MatchHighlighted{}, rp_model.HiddenColumns("name", "path_lower"))
}