package setting

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"strings"
)

type UpdateSetting struct {
	Path        string `json:"path"`
	SyncSetting string `json:"sync_setting"`
}

func (z UpdateSetting) Split() (teamFolder, path string) {
	p := z.Path
	if strings.HasPrefix(z.Path, "/") {
		p = z.Path[1:]
	}
	split := strings.Split(p, "/")
	switch len(split) {
	case 0:
		return split[0], ""
	case 1:
		return split[0], ""
	default:
		return split[0], "/" + strings.Join(split[1:], "/")
	}
}

type Update struct {
	Peer                               dbx_conn.ConnScopedTeam
	File                               fd_file.RowFeed
	Updated                            rp_model.TransactionReport
	ErrorInvalidLineTeamFolderNotFound app_msg.Message
	ErrorInvalidLineSyncSetting        app_msg.Message
	ErrorInvalidLinePathNotFound       app_msg.Message
}

func (z *Update) Preset() {
	z.File.SetModel(&UpdateSetting{})
	z.Updated.SetModel(&UpdateSetting{}, &mo_teamfolder.TeamFolder{})
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeTeamDataContentRead,
		dbx_auth.ScopeTeamDataContentWrite,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
	)
}

func (z *Update) findPath(teamFolder, path string, svt sv_teamfolder.TeamFolder, client dbx_client.Client) (*mo_teamfolder.TeamFolder, *mo_file.Folder, error) {
	tf, err := svt.ResolveByName(teamFolder)
	if err != nil {
		return nil, nil, err
	}
	if path == "" {
		return tf, nil, nil
	}

	client = client.WithPath(dbx_client.Namespace(tf.TeamFolderId))
	f, err := sv_file.NewFiles(client).Resolve(mo_path.NewDropboxPath(path))
	if err != nil {
		return tf, nil, err
	}

	if folder, ok := f.Folder(); ok {
		return tf, folder, nil
	}

	return nil, nil, errors.New("not a folder")
}

func (z *Update) Exec(c app_control.Control) error {
	teamFolders, err := sv_teamfolder.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}
	if err := z.Updated.Open(); err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(z.Peer.Client()).Admin()
	if err != nil {
		return err
	}

	svt := sv_teamfolder.NewCached(z.Peer.Client())
	client := z.Peer.Client().AsAdminId(admin.TeamMemberId)

	err = z.File.Validate(func(m interface{}, rowIndex int) (app_msg.Message, error) {
		setting := m.(*UpdateSetting)
		teamFolder, path := setting.Split()
		found := false
		for _, tf := range teamFolders {
			if strings.ToLower(tf.Name) == strings.ToLower(teamFolder) {
				found = true
				break
			}
		}
		if !found {
			return z.ErrorInvalidLineTeamFolderNotFound.With("TeamFolder", teamFolder), errors.New("team folder not found")
		}

		_, _, err := z.findPath(teamFolder, path, svt, client)
		if err != nil {
			return z.ErrorInvalidLinePathNotFound.With("Error", err).With("Path", path).With("TeamFolder", teamFolder), err
		}
		switch strings.TrimSpace(setting.SyncSetting) {
		case "not_synced", "default":
		// ok
		default:
			return z.ErrorInvalidLineSyncSetting.With("SyncSetting", setting.SyncSetting), errors.New("invalid sync setting")
		}
		return nil, nil
	})
	if err != nil {
		return err
	}

	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		setting := m.(*UpdateSetting)
		teamFolder, path := setting.Split()
		tf, folder, err := z.findPath(teamFolder, path, svt, client)
		if err != nil {
			return err
		}
		var updated *mo_teamfolder.TeamFolder
		if folder != nil {
			updated, err = svt.UpdateSyncSetting(tf, sv_teamfolder.AddNestedSetting(folder, strings.TrimSpace(setting.SyncSetting)))
		} else {
			updated, err = svt.UpdateSyncSetting(tf, sv_teamfolder.RootSyncSetting(strings.TrimSpace(setting.SyncSetting)))
		}
		if err != nil {
			z.Updated.Failure(err, setting)
			return err
		}
		z.Updated.Success(setting, updated)
		return nil
	})
}

func (z *Update) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("update", "/Sales/Report,not_synced")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecMock(c, &Update{}, func(r rc_recipe.Recipe) {
		m := r.(*Update)
		m.File.SetFilePath(f)
	})
}
