package folder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/ingredient/ig_dropbox/ig_file"
	"github.com/watermint/toolbox/ingredient/ig_dropbox/ig_teamspace"
)

type Delete struct {
	Peer           dbx_conn.ConnScopedTeam
	Name           string
	ProgressDelete app_msg.Message
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeAccountInfoRead,
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeTeamDataContentRead,
		dbx_auth.ScopeTeamDataContentWrite,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
}

func (z *Delete) Exec(c app_control.Control) error {
	client, err := ig_teamspace.ClientForRootNamespaceAsAdmin(z.Peer.Client())
	if err != nil {
		return err
	}

	return ig_file.DeleteRecursively(client, mo_path.NewDropboxPath("/"+z.Name), func(path mo_path.DropboxPath) {
		c.UI().Progress(z.ProgressDelete.With("Path", path.Path()))
	})
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Name = "Test"
	})
}
