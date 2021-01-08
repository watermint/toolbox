package sv_teamfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
	"sync"
)

type TeamFolder interface {
	// Retrieve all team folders of the team.
	List() (teamfolders []*mo_teamfolder.TeamFolder, err error)

	// Resolve team folder with team_folder_id.
	// Please use IsNotFound to make sure if an error was not_found error or not.
	Resolve(teamFolderId string) (teamfolder *mo_teamfolder.TeamFolder, err error)

	// Resolve team folder with team folder name.
	// Please use IsNotFound to make sure if an error was not_found error or not.
	ResolveByName(teamFolderName string) (teamfolder *mo_teamfolder.TeamFolder, err error)

	// Create a team folder.
	Create(name string, opts ...CreateOption) (teamfolder *mo_teamfolder.TeamFolder, err error)

	// Activate a team folder from archive.
	Activate(tf *mo_teamfolder.TeamFolder) (teamfolder *mo_teamfolder.TeamFolder, err error)

	// Archive a team folder.
	Archive(tf *mo_teamfolder.TeamFolder) (teamfolder *mo_teamfolder.TeamFolder, err error)

	// Rename the team folder.
	Rename(tf *mo_teamfolder.TeamFolder, newName string) (updated *mo_teamfolder.TeamFolder, err error)

	// Permanently delete archived team folder.
	PermDelete(tf *mo_teamfolder.TeamFolder) (err error)

	// Update sync setting
	UpdateSyncSetting(tf *mo_teamfolder.TeamFolder, opts ...SyncSettingOpt) (teamfolder *mo_teamfolder.TeamFolder, err error)
}

const (
	SyncSettingSkipSetting = ""
	SyncSettingDefault     = "default"
	SyncSettingNotSynced   = "not_synced"
)

func IsNotFound(err error) bool {
	switch err {
	case nil:
		return false
	case ErrorTeamFolderNotFound:
		return true
	}
	if dbx_error.NewErrors(err).IsIdNotFound() {
		return true
	}
	return false
}

var (
	ErrorTeamFolderNotFound = errors.New("team folder not found")
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

func NewCached(ctx dbx_context.Context) TeamFolder {
	return &teamFolderCached{
		impl:  New(ctx),
		cache: nil,
	}
}

type teamFolderCached struct {
	impl  TeamFolder
	cache []*mo_teamfolder.TeamFolder
	mutex sync.Mutex
}

func (z *teamFolderCached) List() (teamfolders []*mo_teamfolder.TeamFolder, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	if z.cache == nil {
		z.cache, err = z.impl.List()
		if err != nil {
			return nil, err
		}
	}
	return z.cache, nil
}

func (z *teamFolderCached) Resolve(teamFolderId string) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	if z.cache == nil {
		z.cache, err = z.impl.List()
		if err != nil {
			return nil, err
		}
	}
	for _, f := range z.cache {
		if f.TeamFolderId == teamFolderId {
			return f, nil
		}
	}
	return nil, ErrorTeamFolderNotFound
}

func (z *teamFolderCached) ResolveByName(teamFolderName string) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	if z.cache == nil {
		z.cache, err = z.impl.List()
		if err != nil {
			return nil, err
		}
	}
	for _, f := range z.cache {
		if f.Name == teamFolderName {
			return f, nil
		}
	}
	return nil, ErrorTeamFolderNotFound
}

func (z *teamFolderCached) Create(name string, opts ...CreateOption) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	z.cache = nil // invalidate cache
	return z.impl.Create(name, opts...)
}

func (z *teamFolderCached) Activate(tf *mo_teamfolder.TeamFolder) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	z.cache = nil // invalidate cache
	return z.impl.Activate(tf)
}

func (z *teamFolderCached) Archive(tf *mo_teamfolder.TeamFolder) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	z.cache = nil // invalidate cache
	return z.impl.Archive(tf)
}

func (z *teamFolderCached) Rename(tf *mo_teamfolder.TeamFolder, newName string) (updated *mo_teamfolder.TeamFolder, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	z.cache = nil // invalidate cache
	return z.impl.Rename(tf, newName)
}

func (z *teamFolderCached) PermDelete(tf *mo_teamfolder.TeamFolder) (err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	z.cache = nil // invalidate cache
	return z.impl.PermDelete(tf)
}

func (z *teamFolderCached) UpdateSyncSetting(tf *mo_teamfolder.TeamFolder, opts ...SyncSettingOpt) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	z.cache = nil // invalidate cache
	return z.impl.UpdateSyncSetting(tf, opts...)
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

func (z *teamFolderImpl) ResolveByName(teamFolderName string) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	teamfolders, err := z.List()
	if err != nil {
		return nil, err
	}
	for _, f := range teamfolders {
		if f.Name == teamFolderName {
			return f, nil
		}
	}
	return nil, ErrorTeamFolderNotFound
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
