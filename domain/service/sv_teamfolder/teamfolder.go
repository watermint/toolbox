package sv_teamfolder

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
)

type TeamFolder interface {
	List() (teamfolders []*mo_teamfolder.TeamFolder, err error)
	Resolve(teamFolderId string) (teamfolder *mo_teamfolder.TeamFolder, err error)
	Create(name string, opts ...CreateOption) (teamfolder *mo_teamfolder.TeamFolder, err error)
	Activate(tf *mo_teamfolder.TeamFolder) (err error)
	Archive(tf *mo_teamfolder.TeamFolder) (err error)
	Rename(tf *mo_teamfolder.TeamFolder, newName string) (updated *mo_teamfolder.TeamFolder, err error)
	PermDelete(tf *mo_teamfolder.TeamFolder) (err error)
}

type createOptions struct {
}

type CreateOption func(opt *createOptions) *createOptions

func New(ctx api_context.Context) TeamFolder {
	return &teamFolderImpl{
		ctx: ctx,
	}
}

type teamFolderImpl struct {
	ctx api_context.Context
}

func (z *teamFolderImpl) List() (teamfolders []*mo_teamfolder.TeamFolder, err error) {
	panic("implement me")
}

func (z *teamFolderImpl) Resolve(teamFolderId string) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	panic("implement me")
}

func (z *teamFolderImpl) Create(name string, opts ...CreateOption) (teamfolder *mo_teamfolder.TeamFolder, err error) {
	panic("implement me")
}

func (z *teamFolderImpl) Activate(tf *mo_teamfolder.TeamFolder) (err error) {
	panic("implement me")
}

func (z *teamFolderImpl) Archive(tf *mo_teamfolder.TeamFolder) (err error) {
	panic("implement me")
}

func (z *teamFolderImpl) Rename(tf *mo_teamfolder.TeamFolder, newName string) (updated *mo_teamfolder.TeamFolder, err error) {
	panic("implement me")
}

func (z *teamFolderImpl) PermDelete(tf *mo_teamfolder.TeamFolder) (err error) {
	panic("implement me")
}
