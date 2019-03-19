package sv_sharedfolder_mount

import "github.com/watermint/toolbox/domain/model/mo_sharedfolder"

type SharedFolderMount interface {
	List() (mount []*mo_sharedfolder.SharedFolder, err error)
	Mount(sf *mo_sharedfolder.SharedFolder) (mount *mo_sharedfolder.SharedFolder, err error)
	Unmount(sf *mo_sharedfolder.SharedFolder) (err error)
}
