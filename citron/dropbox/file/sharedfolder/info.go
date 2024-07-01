package sharedfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Info struct {
	Peer           dbx_conn.ConnScopedIndividual
	SharedFolderId string
	Policies       rp_model.RowReport
}

func (z *Info) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeSharingRead,
	)
	z.Policies.SetModel(
		&mo_sharedfolder.SharedFolder{},
		rp_model.HiddenColumns(
			"parent_shared_folder_id",
			"owner_team_id",
		),
	)
}

func (z *Info) Exec(c app_control.Control) error {
	c.Log().Debug("Scanning folders")
	folder, err := sv_sharedfolder.New(z.Peer.Client()).Resolve(z.SharedFolderId)
	if err != nil {
		return err
	}

	if err := z.Policies.Open(); err != nil {
		return err
	}

	z.Policies.Row(folder)
	return nil
}

func (z *Info) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Info{}, func(r rc_recipe.Recipe) {
		m := r.(*Info)
		m.SharedFolderId = "1234567890"
	})
}
