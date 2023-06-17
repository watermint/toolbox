package setting

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type ScanTarget struct {
	TeamFolder   *mo_teamfolder.TeamFolder `json:"team_folder"`
	Path         string                    `json:"path,omitempty"`
	CurrentDepth int                       `json:"current_depth"`
	ScanDepth    int                       `json:"scan_depth"`
}

type FolderSettings struct {
	TeamFolder  string `json:"team_folder"`
	Path        string `json:"path"`
	SyncSetting string `json:"sync_setting"`
}

const (
	queueScanTeamFolder = "scan_folder"
)

type List struct {
	Peer       dbx_conn.ConnScopedTeam
	SubFolders kv_storage.Storage
	Folders    rp_model.RowReport
	Settings   rp_model.RowReport
	ScanAll    bool
	ShowAll    bool
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeTeamDataContentRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
	)
	z.Folders.SetModel(&mo_file.ConcreteEntry{})
	z.Settings.SetModel(&FolderSettings{})
}

func (z *List) scanTeamFolder(st *ScanTarget, s eq_sequence.Stage, client dbx_client.Client) error {
	if st.ScanDepth < st.CurrentDepth {
		return nil
	}

	q := s.Get(queueScanTeamFolder)
	c := client.WithPath(dbx_client.Namespace(st.TeamFolder.TeamFolderId))

	entries, err := sv_file.NewFiles(c).List(mo_path.NewDropboxPath(st.Path),
		sv_file.IncludeMountedFolders(true),
		sv_file.Recursive(false))
	if err != nil {
		return err
	}

	settings := make(map[string]string)
	for _, setting := range st.TeamFolder.ContentSyncSettings() {
		settings[setting.Id] = setting.SyncSetting
	}

	for _, entry := range entries {
		if folder, isFolder := entry.Folder(); isFolder {
			if setting, ok := settings[folder.Id]; ok {
				z.Settings.Row(&FolderSettings{
					TeamFolder:  st.TeamFolder.Name,
					Path:        folder.PathDisplay(),
					SyncSetting: setting,
				})
			} else if z.ShowAll {
				z.Settings.Row(&FolderSettings{
					TeamFolder:  st.TeamFolder.Name,
					Path:        folder.PathDisplay(),
					SyncSetting: "",
				})
			}
			q.Enqueue(&ScanTarget{
				TeamFolder:   st.TeamFolder,
				Path:         folder.PathDisplay(),
				CurrentDepth: st.CurrentDepth + 1,
				ScanDepth:    st.ScanDepth,
			})
		}
	}
	return nil
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Folders.Open(); err != nil {
		return err
	}
	if err := z.Settings.Open(); err != nil {
		return err
	}

	teamfolders, err := sv_teamfolder.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(z.Peer.Client()).Admin()
	if err != nil {
		return err
	}

	client := z.Peer.Client().AsAdminId(admin.TeamMemberId)

	features, err := sv_team.New(z.Peer.Client()).Feature()
	if err != nil {
		return err
	}
	scanDepth := 2
	if features.HasTeamSharedDropbox {
		scanDepth = 3
	}
	if z.ScanAll {
		scanDepth = 1000
	}

	for _, teamfolder := range teamfolders {
		z.Settings.Row(&FolderSettings{
			TeamFolder:  teamfolder.Name,
			Path:        "",
			SyncSetting: teamfolder.SyncSetting,
		})
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define(queueScanTeamFolder, z.scanTeamFolder, s, client)
		q := s.Get(queueScanTeamFolder)

		for _, teamfolder := range teamfolders {
			q.Enqueue(&ScanTarget{
				TeamFolder:   teamfolder,
				Path:         "",
				CurrentDepth: 0,
				ScanDepth:    scanDepth,
			})
		}
	})

	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
