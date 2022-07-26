package sharedfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Leave struct {
	Peer           dbx_conn.ConnScopedIndividual
	SharedFolderId string
	KeepCopy       bool
}

func (z *Leave) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
	)
}

func (z *Leave) Exec(c app_control.Control) error {
	err := sv_sharedfolder.New(z.Peer.Context()).Leave(&mo_sharedfolder.SharedFolder{
		SharedFolderId: z.SharedFolderId,
	}, sv_sharedfolder.LeaveACopy(z.KeepCopy))
	return err
}

func (z *Leave) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Leave{}, func(r rc_recipe.Recipe) {
		m := r.(*Leave)
		m.SharedFolderId = "123456"
	})
}
