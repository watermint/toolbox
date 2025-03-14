package mount

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Delete struct {
	Peer                 dbx_conn.ConnScopedIndividual
	SharedFolderId       string
	Mount                rp_model.RowReport
	InfoAlreadyUnmounted app_msg.Message
	BasePath             mo_string.SelectString
}

func (z *Delete) Preset() {
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
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Delete) Exec(c app_control.Control) error {
	if err := z.Mount.Open(); err != nil {
		return err
	}

	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	err := sv_sharedfolder_mount.New(client).Unmount(&mo_sharedfolder.SharedFolder{SharedFolderId: z.SharedFolderId})
	if err != nil {
		de := dbx_error.NewErrors(err)
		switch {
		case de.HasPrefix("access_error/unmounted"):
			c.UI().Info(z.InfoAlreadyUnmounted)

		default:
			return err
		}
	}

	mount, err := sv_sharedfolder.New(client).Resolve(z.SharedFolderId)
	if err != nil {
		return err
	}
	z.Mount.Row(mount)

	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.SharedFolderId = "123456"
	})
}
