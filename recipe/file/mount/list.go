package mount

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer  dbx_conn.ConnScopedIndividual
	Mount rp_model.RowReport
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeSharingRead,
	)
	z.Mount.SetModel(&mo_sharedfolder.SharedFolder{},
		rp_model.HiddenColumns(
			"owner_team_id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Mount.Open(); err != nil {
		return err
	}

	mounts, err := sv_sharedfolder_mount.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	for _, mount := range mounts {
		z.Mount.Row(mount)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
