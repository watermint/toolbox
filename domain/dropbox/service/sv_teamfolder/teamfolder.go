package sv_teamfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type TeamFolder interface {
	List() (teamfolders []*mo_teamfolder.TeamFolder, err error)
	Resolve(teamFolderId string) (teamfolder *mo_teamfolder.TeamFolder, err error)
	Create(name string, opts ...CreateOption) (teamfolder *mo_teamfolder.TeamFolder, err error)
	Activate(tf *mo_teamfolder.TeamFolder) (teamfolder *mo_teamfolder.TeamFolder, err error)
	Archive(tf *mo_teamfolder.TeamFolder) (teamfolder *mo_teamfolder.TeamFolder, err error)
	Rename(tf *mo_teamfolder.TeamFolder, newName string) (updated *mo_teamfolder.TeamFolder, err error)
	PermDelete(tf *mo_teamfolder.TeamFolder) (err error)
	UpdateSyncSetting(tf *mo_teamfolder.TeamFolder, opts ...SyncSettingOpt) (teamfolder *mo_teamfolder.TeamFolder, err error)
}

const (
	SyncSettingSkipSetting = ""
	SyncSettingDefault     = "default"
	SyncSettingNotSynced   = "not_synced"
)

type SyncSettingsOpts struct {
	TeamFolderId        string              `json:"team_folder_id"`
	SyncSetting         string              `json:"sync_setting,omitempty"`
	ContentSyncSettings []NestedSyncSetting `json:"content_sync_settings,omitempty"`
}

func (z SyncSettingsOpts) Apply(opts []SyncSettingOpt) SyncSettingsOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

type SyncSettingOpt func(o SyncSettingsOpts) SyncSettingsOpts

func RootSyncSetting(setting string) SyncSettingOpt {
	return func(o SyncSettingsOpts) SyncSettingsOpts {
		o.SyncSetting = setting
		return o
	}
}

func AddNestedSetting(folder mo_file.Entry, setting string) SyncSettingOpt {
	return func(o SyncSettingsOpts) SyncSettingsOpts {
		o.ContentSyncSettings = append(o.ContentSyncSettings, NestedSyncSetting{
			FolderId:    folder.Concrete().Id,
			SyncSetting: setting,
		})
		return o
	}
}

type NestedSyncSetting struct {
	FolderId    string `json:"id"`
	SyncSetting string `json:"sync_setting"`
}

type createOptions struct {
	syncSetting string
}

type CreateOption func(opt *createOptions) *createOptions

func SyncDefault() CreateOption {
	return func(opt *createOptions) *createOptions {
		opt.syncSetting = SyncSettingDefault
		return opt
	}
}
func SyncNoSync() CreateOption {
	return func(opt *createOptions) *createOptions {
		opt.syncSetting = SyncSettingNotSynced
		return opt
	}
}

func New(ctx dbx_context.Context) TeamFolder {
	return &teamFolderImpl{
		ctx: ctx,
	}
}

type teamFolderImpl struct {
	ctx dbx_context.Context
}

func (z *teamFolderImpl) UpdateSyncSetting(tf *mo_teamfolder.TeamFolder, opts ...SyncSettingOpt) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	ss := SyncSettingsOpts{}.Apply(opts)
	ss.TeamFolderId = tf.TeamFolderId

	res := z.ctx.Post("team/team_folder/update_sync_settings", api_request.Param(&ss))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	teamfolder = &mo_teamfolder.TeamFolder{}
	err = res.Success().Json().Model(teamfolder)
	return
}

func (z *teamFolderImpl) List() (teamfolders []*mo_teamfolder.TeamFolder, err error) {
	teamfolders = make([]*mo_teamfolder.TeamFolder, 0)
	res := z.ctx.List("team/team_folder/list").Call(
		dbx_list.Continue("team/team_folder/list/continue"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("team_folders"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			tf := &mo_teamfolder.TeamFolder{}
			if err := entry.Model(tf); err != nil {
				return err
			}
			teamfolders = append(teamfolders, tf)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return teamfolders, nil
}

func (z *teamFolderImpl) Resolve(teamFolderId string) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	p := struct {
		TeamFolderIds []string `json:"team_folder_ids"`
	}{
		TeamFolderIds: []string{teamFolderId},
	}
	res := z.ctx.Post("team/team_folder/get_info", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	teamfolder = &mo_teamfolder.TeamFolder{}
	err = res.Success().Json().FindModel(es_json.PathArrayFirst, teamfolder)
	return
}

func (z *teamFolderImpl) Create(name string, opts ...CreateOption) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	co := &createOptions{}
	for _, o := range opts {
		o(co)
	}
	p := struct {
		Name        string `json:"name"`
		SyncSetting string `json:"sync_setting,omitempty"`
	}{
		Name:        name,
		SyncSetting: co.syncSetting,
	}

	res := z.ctx.Post("team/team_folder/create", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	teamfolder = &mo_teamfolder.TeamFolder{}
	err = res.Success().Json().Model(teamfolder)
	return
}

func (z *teamFolderImpl) Activate(tf *mo_teamfolder.TeamFolder) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	p := struct {
		TeamFolderId string `json:"team_folder_id"`
	}{
		TeamFolderId: tf.TeamFolderId,
	}
	res := z.ctx.Post("team/team_folder/activate", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	teamfolder = &mo_teamfolder.TeamFolder{}
	err = res.Success().Json().Model(teamfolder)
	return
}

func (z *teamFolderImpl) Archive(tf *mo_teamfolder.TeamFolder) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	p := struct {
		TeamFolderId string `json:"team_folder_id"`
	}{
		TeamFolderId: tf.TeamFolderId,
	}
	res := z.ctx.Async("team/team_folder/archive", api_request.Param(p)).Call(
		dbx_async.Status("team/team_folder/archive/check"))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	teamfolder = &mo_teamfolder.TeamFolder{}
	err = res.Success().Json().Model(teamfolder)
	return
}

func (z *teamFolderImpl) Rename(tf *mo_teamfolder.TeamFolder, newName string) (updated *mo_teamfolder.TeamFolder, err error) {
	p := struct {
		TeamFolderId string `json:"team_folder_id"`
		Name         string `json:"name"`
	}{
		TeamFolderId: tf.TeamFolderId,
		Name:         newName,
	}
	res := z.ctx.Post("team/team_folder/rename", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	updated = &mo_teamfolder.TeamFolder{}
	err = res.Success().Json().Model(updated)
	return
}

func (z *teamFolderImpl) PermDelete(tf *mo_teamfolder.TeamFolder) (err error) {
	p := struct {
		TeamFolderId string `json:"team_folder_id"`
	}{
		TeamFolderId: tf.TeamFolderId,
	}
	res := z.ctx.Post("team/team_folder/permanently_delete", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}
