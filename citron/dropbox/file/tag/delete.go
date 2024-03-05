package tag

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_tag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Delete struct {
	Peer dbx_conn.ConnScopedIndividual
	Path mo_path.DropboxPath
	Tag  string
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeFilesMetadataWrite,
	)
}

func (z *Delete) Exec(c app_control.Control) error {
	return sv_file_tag.New(z.Peer.Client()).Delete(z.Path, z.Tag)
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Tag = "delete"
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("delete")
	})
}
