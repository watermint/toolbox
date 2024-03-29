package lock

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_lock"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type List struct {
	Peer                       dbx_conn.ConnScopedTeam
	TeamFolder                 string
	Path                       mo_path.DropboxPath
	Lock                       rp_model.RowReport
	ErrorTeamSpaceNotSupported app_msg.Message
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.Lock.SetModel(
		&mo_file.LockInfo{},
		rp_model.HiddenColumns(
			"id",
			"path_lower",
			"revision",
			"content_hash",
			"shared_folder_id",
			"parent_shared_folder_id",
			"lock_holder_account_id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if ok, _ := sv_team.UnlessTeamFolderApiSupported(z.Peer.Client()); ok {
		c.UI().Error(z.ErrorTeamSpaceNotSupported)
		return errors.New("team space is not supported by this command")
	}

	if err := z.Lock.Open(); err != nil {
		return err
	}

	teamFolder, err := sv_teamfolder.New(z.Peer.Client()).ResolveByName(z.TeamFolder)
	if err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(z.Peer.Client()).Admin()
	if err != nil {
		return err
	}

	ctx := z.Peer.Client().WithPath(dbx_client.Namespace(teamFolder.TeamFolderId)).AsAdminId(admin.TeamMemberId)

	return sv_file_lock.New(ctx).List(z.Path, func(entry *mo_file.LockInfo) {
		z.Lock.Row(entry)
	})
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.Path = mo_path.NewDropboxPath("/")
		m.TeamFolder = qtr_endtoend.TestTeamFolderName
	})
}
