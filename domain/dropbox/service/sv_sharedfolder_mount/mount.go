package sv_sharedfolder_mount

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Mount interface {
	List() (mount []*mo_sharedfolder.SharedFolder, err error)
	Mount(sf *mo_sharedfolder.SharedFolder) (mount *mo_sharedfolder.SharedFolder, err error)
	Unmount(sf *mo_sharedfolder.SharedFolder) (err error)
}

func New(ctx dbx_context.Context) Mount {
	return &mountImpl{
		ctx: ctx,
	}
}

type mountImpl struct {
	ctx dbx_context.Context
}

func (z *mountImpl) List() (mount []*mo_sharedfolder.SharedFolder, err error) {
	mount = make([]*mo_sharedfolder.SharedFolder, 0)
	res := z.ctx.List("sharing/list_mountable_folders").Call(
		dbx_list.Continue("sharing/list_mountable_folders/continue"),
		dbx_list.ResultTag("entries"),
		dbx_list.OnEntry(func(entry tjson.Json) error {
			m := &mo_sharedfolder.SharedFolder{}
			if err = entry.Model(m); err != nil {
				return err
			}
			mount = append(mount, m)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
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

	res := z.ctx.Post("sharing/mount_folder", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	mount = &mo_sharedfolder.SharedFolder{}
	err = res.Success().Json().Model(mount)
	return
}

func (z *mountImpl) Unmount(sf *mo_sharedfolder.SharedFolder) (err error) {
	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
	}{
		SharedFolderId: sf.SharedFolderId,
	}
	res := z.ctx.Post("sharing/unmount_folder", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}
