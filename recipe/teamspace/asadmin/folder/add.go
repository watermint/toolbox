package folder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_folder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/ingredient/ig_dropbox/ig_teamspace"
)

type Add struct {
	Peer    dbx_conn.ConnScopedTeam
	Name    string
	Created rp_model.RowReport
}

func (z *Add) Preset() {
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
	z.Created.SetModel(
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"id",
			"tag",
			"path_lower",
			"client_modified",
			"server_modified",
			"revision",
			"size",
			"content_hash",
		),
	)
}

func (z *Add) Exec(c app_control.Control) error {
	client, err := ig_teamspace.ClientForRootNamespaceAsAdmin(z.Peer.Client())
	if err != nil {
		return err
	}

	if err := z.Created.Open(); err != nil {
		return err
	}

	entry, err := sv_file_folder.New(client).Create(mo_path.NewDropboxPath("/" + z.Name))
	if err != nil {
		return err
	}

	z.Created.Row(entry.Concrete())
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Name = "Test"
	})
}
