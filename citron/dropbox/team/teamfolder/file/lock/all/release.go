package all

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
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type PathLock struct {
	Path string `json:"path"`
}

type Release struct {
	Peer                       dbx_conn.ConnScopedTeam
	TeamFolder                 string
	Path                       mo_path.DropboxPath
	BatchSize                  int
	OperationLog               rp_model.TransactionReport
	ErrorTeamSpaceNotSupported app_msg.Message
}

func (z *Release) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.BatchSize = 100
	z.OperationLog.SetModel(
		&PathLock{},
		&mo_file.LockInfo{},
		rp_model.HiddenColumns(
			"result.id",
			"result.name",
			"result.path_lower",
			"result.path_display",
			"result.revision",
			"result.content_hash",
			"result.shared_folder_id",
			"result.parent_shared_folder_id",
			"result.lock_holder_account_id",
		),
	)
}

func (z *Release) Exec(c app_control.Control) error {
	if ok, _ := sv_team.UnlessTeamFolderApiSupported(z.Peer.Client()); ok {
		c.UI().Error(z.ErrorTeamSpaceNotSupported)
		return errors.New("team space is not supported by this command")
	}

	if err := z.OperationLog.Open(); err != nil {
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
	l := c.Log()
	sfl := sv_file_lock.New(ctx)

	var lastErr error

	unlockBucket := func(bucket []mo_path.DropboxPath) {
		if len(bucket) < 1 {
			l.Debug("No more bucket")
			return
		}

		unlocked, err := sfl.UnlockBatch(bucket)
		if err != nil {
			l.Debug("Something happened during unlock batch", esl.Error(err))
			lastErr = err
		}
		for path, info := range unlocked {
			if info.Error != nil {
				z.OperationLog.Failure(info.Error, PathLock{
					Path: path,
				})
			} else {
				z.OperationLog.Success(PathLock{
					Path: path,
				}, info.Entry.LockInfo())
			}
		}
	}

	bucket := make([]mo_path.DropboxPath, 0)
	listErr := sfl.List(z.Path, func(entry *mo_file.LockInfo) {
		bucket = append(bucket, mo_path.NewDropboxPath(entry.PathDisplay))
		if z.BatchSize <= len(bucket) {
			l.Debug("Bucket exceeds batch size, release those", esl.Int("bucketLen", len(bucket)))
			unlockBucket(bucket)
			bucket = make([]mo_path.DropboxPath, 0)
		}
	})
	unlockBucket(bucket)
	if listErr != nil {
		return listErr
	}

	return lastErr
}

func (z *Release) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Release{}, func(r rc_recipe.Recipe) {
		m := r.(*Release)
		m.Path = mo_path.NewDropboxPath("/")
		m.TeamFolder = qtr_endtoend.TestTeamFolderName
	})
}
