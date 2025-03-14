package _import

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_url"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Url struct {
	rc_recipe.RemarkIrreversible
	Peer           dbx_conn.ConnScopedIndividual
	Path           mo_path.DropboxPath
	Url            string
	OperationLog   rp_model.RowReport
	ProgressImport app_msg.Message
	BasePath       mo_string.SelectString
}

func (z *Url) Preset() {
	z.Peer.SetScopes(dbx_auth.ScopeFilesContentWrite)
	z.OperationLog.SetModel(
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"id",
			"path_lower",
			"content_hash",
			"shared_folder_id",
			"parent_shared_folder_id",
		),
	)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Url) Exec(c app_control.Control) error {
	ui := c.UI()
	ctx := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	path := sv_file_url.PathWithName(z.Path, z.Url)
	ui.Progress(z.ProgressImport.With("Path", path.Path()).With("Url", z.Url))
	entry, err := sv_file_url.New(ctx).Save(path, z.Url)
	if err != nil {
		return err
	}
	z.OperationLog.Row(entry.Concrete())
	return nil
}

func (z *Url) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Url{}, func(r rc_recipe.Recipe) {
		ru := r.(*Url)
		ru.Url = "https://dummyimage.com/10x10/000/fff"
		ru.Path = qtr_endtoend.NewTestDropboxFolderPath("file-import-url")
	})
}
