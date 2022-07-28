package mount

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Add struct {
	Peer               dbx_conn.ConnScopedIndividual
	SharedFolderId     string
	Mount              rp_model.RowReport
	InfoAlreadyMounted app_msg.Message
}

func (z *Add) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
	)
	z.Mount.SetModel(
		&mo_sharedfolder.SharedFolder{},
		rp_model.HiddenColumns(
			"parent_shared_folder_id",
			"team_member_id",
			"namespace_id",
			"owner_team_id",
		),
	)
}

func (z *Add) Exec(c app_control.Control) error {
	if err := z.Mount.Open(); err != nil {
		return err
	}

	mount, err := sv_sharedfolder_mount.New(z.Peer.Context()).Mount(&mo_sharedfolder.SharedFolder{SharedFolderId: z.SharedFolderId})
	if err != nil {
		de := dbx_error.NewErrors(err)
		switch {
		case de.HasPrefix("already_mounted"):
			c.UI().Info(z.InfoAlreadyMounted)

		default:
			return err
		}
	}

	z.Mount.Row(mount)
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.SharedFolderId = "123456"
	})
}
