package revision

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	mo_path2 "github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"os"
)

type Download struct {
	Peer      dbx_conn.ConnScopedIndividual
	Revision  string
	LocalPath mo_path.FileSystemPath
	Entry     rp_model.RowReport
}

func (z *Download) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesMetadataRead,
	)
	z.Entry.SetModel(
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"tag",
			"path_lower",
			"shared_folder_id",
			"parent_shared_folder_id",
		),
	)
}

func (z *Download) Exec(c app_control.Control) error {
	if err := z.Entry.Open(); err != nil {
		return err
	}
	svd := sv_file_content.NewDownload(z.Peer.Context())
	entry, path, err := svd.Download(mo_path2.NewDropboxPath("rev:" + z.Revision))
	if err != nil {
		return err
	}
	err = os.Rename(path.Path(), z.LocalPath.Path())
	if err != nil {
		return err
	}
	z.Entry.Row(entry)
	return nil
}

func (z *Download) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Download{}, func(r rc_recipe.Recipe) {
		m := r.(*Download)
		m.Revision = "a1c10ce0dd78"
		m.LocalPath = qtr_endtoend.NewTestFileSystemFolderPath(c, "download")
	})
}
