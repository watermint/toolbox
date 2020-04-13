package search

import (
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Name struct {
	Peer      dbx_conn.ConnUserFile
	Path      mo_string.OptionalString
	Query     string
	Extension mo_string.OptionalString
	Category  mo_string.SelectString
	Matches   rp_model.RowReport
}

func (z *Name) Exec(c app_control.Control) error {
	so := make([]sv_file.SearchOpt, 0)
	so = append(so, sv_file.SearchFileNameOnly())
	so = append(so, sv_file.SearchIncludeHighlights())
	if z.Extension.IsExists() {
		so = append(so, sv_file.SearchFileExtension(z.Extension.Value()))
	}
	if z.Category.Value() != "" {
		so = append(so, sv_file.SearchCategories(z.Category.Value()))
	}
	if z.Path.IsExists() {
		so = append(so, sv_file.SearchPath(mo_path.NewDropboxPath(z.Path.Value())))
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
	z.Category.SetOptions(
		[]string{"", "image", "document", "pdf", "spreadsheet", "presentation", "audio", "video", "folder", "paper", "others"},
		"",
	)
	z.Matches.SetModel(&mo_file.MatchHighlighted{},
		rp_model.HiddenColumns(
			"name",
			"path_lower"),
	)
}
