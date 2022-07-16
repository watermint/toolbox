package revision

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_revision"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type List struct {
	Peer      dbx_conn.ConnScopedIndividual
	Path      mo_path.DropboxPath
	Revisions rp_model.RowReport
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesMetadataRead,
	)
	z.Revisions.SetModel(&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"tag",
			"path_lower",
			"shared_folder_id",
			"parent_shared_folder_id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Revisions.Open(); err != nil {
		return err
	}

	svr := sv_file_revision.New(z.Peer.Context())
	revisions, err := svr.List(z.Path)
	if err != nil {
		return err
	}
	for _, rev := range revisions.Entries {
		z.Revisions.Row(rev)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("/rev/list.txt")
	})
}
