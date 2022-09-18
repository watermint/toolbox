package folder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/teamspace"
)

type Permdelete struct {
	Peer dbx_conn.ConnScopedTeam
	Name string
}

func (z *Permdelete) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeAccountInfoRead,
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeFilesPermanentDelete,
		dbx_auth.ScopeTeamDataContentRead,
		dbx_auth.ScopeTeamDataContentWrite,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
}

func (z *Permdelete) Exec(c app_control.Control) error {
	client, err := teamspace.ClientForRootNamespaceAsAdmin(z.Peer.Client())
	if err != nil {
		return err
	}

	return sv_file.NewFiles(client).PermDelete(mo_path.NewDropboxPath("/" + z.Name))
}

func (z *Permdelete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Permdelete{}, func(r rc_recipe.Recipe) {
		m := r.(*Permdelete)
		m.Name = "Test"
	})
}
