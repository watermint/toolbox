package revision

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_restore"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Restore struct {
	Peer     dbx_conn.ConnScopedIndividual
	Path     mo_path.DropboxPath
	Revision string
	Entry    rp_model.RowReport
}

func (z *Restore) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
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

func (z *Restore) Exec(c app_control.Control) error {
	if err := z.Entry.Open(); err != nil {
		return err
	}
	svr := sv_file_restore.New(z.Peer.Context())
	entry, err := svr.Restore(z.Path, z.Revision)

	if err != nil {
		return err
	}
	z.Entry.Row(entry)
	return nil
}

func (z *Restore) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Restore{}, func(r rc_recipe.Recipe) {
		m := r.(*Restore)
		m.Path = mo_path.NewDropboxPath("/root/word.docx")
		m.Revision = "a1c10ce0dd78"
	})
}
