package sv_sharedfolder_mount

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
)

type Mount interface {
	List() (mount []*mo_sharedfolder.SharedFolder, err error)
	Mount(sf *mo_sharedfolder.SharedFolder) (mount *mo_sharedfolder.SharedFolder, err error)
	Unmount(sf *mo_sharedfolder.SharedFolder) (err error)
}

func New(ctx api_context.Context) Mount {
	return &mountImpl{
		ctx: ctx,
	}
}

type mountImpl struct {
	ctx api_context.Context
}

func (z *mountImpl) List() (mount []*mo_sharedfolder.SharedFolder, err error) {
	mount = make([]*mo_sharedfolder.SharedFolder, 0)
	err = z.ctx.List("sharing/list_mountable_folders").
		Continue("sharing/list_mountable_folders/continue").
		UseHasMore(false).
		ResultTag("entries").
		OnEntry(func(entry api_list.ListEntry) error {
			m := &mo_sharedfolder.SharedFolder{}
			if err = entry.Model(m); err != nil {
				return err
			}
			mount = append(mount, m)
			return nil
		}).
		Call()
	if err != nil {
		return nil, err
	}
	return mount, nil
}

func (z *mountImpl) Mount(sf *mo_sharedfolder.SharedFolder) (mount *mo_sharedfolder.SharedFolder, err error) {
	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
	}{
		SharedFolderId: sf.SharedFolderId,
	}

	mount = &mo_sharedfolder.SharedFolder{}
	res, err := z.ctx.Rpc("sharing/mount_folder").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err = res.Model(mount); err != nil {
		return nil, err
	}
	return mount, nil
}

func (z *mountImpl) Unmount(sf *mo_sharedfolder.SharedFolder) (err error) {
	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
	}{
		SharedFolderId: sf.SharedFolderId,
	}
	_, err = z.ctx.Rpc("sharing/unmount_folder").Param(p).Call()
	if err != nil {
		return err
	}
	return nil
}
