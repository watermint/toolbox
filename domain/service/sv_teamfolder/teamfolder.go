package sv_teamfolder

import "github.com/watermint/toolbox/domain/model/mo_teamfolder"

type TeamFolder interface {
	List() (teamfolders []*mo_teamfolder.TeamFolder, err error)
	Resolve(teamFolderId string) (teamfolder *mo_teamfolder.TeamFolder, err error)
	Create(name string) (teamfolder *mo_teamfolder.TeamFolder, err error)
	Archive()
}
