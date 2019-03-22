package sv_sharedfolder

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_profile"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
)

type SharedFolder interface {
	Create(path mo_path.Path, opts ...CreateOption) (sf *mo_sharedfolder.SharedFolder, err error)
	Delete(sf *mo_sharedfolder.SharedFolder, opts ...DeleteOption) (err error)
	List() (sf []*mo_sharedfolder.SharedFolder, err error)
	Leave(sf *mo_sharedfolder.SharedFolder, opts ...DeleteOption) (err error)
	Resolve(sharedFolderId string) (sf *mo_sharedfolder.SharedFolder, err error)
	Transfer(sf *mo_sharedfolder.SharedFolder, to TransferTo) (err error)
}

func New(ctx api_context.Context) SharedFolder {
	return &sharedFolderImpl{
		ctx: ctx,
	}
}

type transferTo struct {
	dropboxId string
}
type TransferTo func(to *transferTo) *transferTo

func ToProfile(p *mo_profile.Profile) TransferTo {
	return func(to *transferTo) *transferTo {
		to.dropboxId = p.AccountId
		return to
	}
}
func ToTeamMemberId(teamMemberId string) TransferTo {
	return func(to *transferTo) *transferTo {
		to.dropboxId = teamMemberId
		return to
	}
}

type createOptions struct {
}

type CreateOption func(opt *createOptions) *createOptions

type deleteOptions struct {
	leaveACopy bool
}
type DeleteOption func(opt *deleteOptions) *deleteOptions

func LeaveACopy() DeleteOption {
	return func(opt *deleteOptions) *deleteOptions {
		opt.leaveACopy = true
		return opt
	}
}

type sharedFolderImpl struct {
	ctx   api_context.Context
	limit int
}

func (z *sharedFolderImpl) Transfer(sf *mo_sharedfolder.SharedFolder, to TransferTo) (err error) {
	too := &transferTo{}
	to(too)

	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
		ToDropboxId    string `json:"to_dropbox_id"`
	}{
		SharedFolderId: sf.SharedFolderId,
		ToDropboxId:    too.dropboxId,
	}

	_, err = z.ctx.Request("sharing/transfer_folder").Param(p).Call()
	if err != nil {
		return err
	}
	return nil
}

func (z *sharedFolderImpl) Resolve(sharedFolderId string) (sf *mo_sharedfolder.SharedFolder, err error) {
	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
	}{
		SharedFolderId: sharedFolderId,
	}

	sf = &mo_sharedfolder.SharedFolder{}
	res, err := z.ctx.Request("sharing/get_folder_metadata").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err = res.Model(sf); err != nil {
		return nil, err
	}
	return sf, nil
}

func (z *sharedFolderImpl) Leave(sf *mo_sharedfolder.SharedFolder, opts ...DeleteOption) (err error) {
	do := &deleteOptions{}
	for _, o := range opts {
		o(do)
	}
	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
		LeaveACopy     bool   `json:"leave_a_copy,omitempty"`
	}{
		SharedFolderId: sf.SharedFolderId,
		LeaveACopy:     do.leaveACopy,
	}
	sf = &mo_sharedfolder.SharedFolder{}
	_, err = z.ctx.Async("sharing/relinquish_folder_membership").
		Status("sharing/check_share_job_status").
		Param(p).Call()
	if err != nil {
		return err
	}
	return nil
}

func (z *sharedFolderImpl) Create(path mo_path.Path, opts ...CreateOption) (sf *mo_sharedfolder.SharedFolder, err error) {
	co := &createOptions{}
	for _, o := range opts {
		o(co)
	}

	p := struct {
		Path string `json:"path"`
	}{
		Path: path.Path(),
	}

	sf = &mo_sharedfolder.SharedFolder{}
	res, err := z.ctx.Async("sharing/share_folder").
		Status("sharing/check_share_job_status").
		Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err = res.Model(sf); err != nil {
		return nil, err
	}
	return sf, nil
}

func (z *sharedFolderImpl) Delete(sf *mo_sharedfolder.SharedFolder, opts ...DeleteOption) (err error) {
	do := &deleteOptions{}
	for _, o := range opts {
		o(do)
	}
	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
		LeaveACopy     bool   `json:"leave_a_copy,omitempty"`
	}{
		SharedFolderId: sf.SharedFolderId,
		LeaveACopy:     do.leaveACopy,
	}
	sf = &mo_sharedfolder.SharedFolder{}
	_, err = z.ctx.Async("sharing/unshare_folder").
		Status("sharing/check_share_job_status").
		Param(p).Call()
	if err != nil {
		return err
	}
	return nil
}

func (z *sharedFolderImpl) List() (sf []*mo_sharedfolder.SharedFolder, err error) {
	p := struct {
		Limit int `json:"limit,omitempty"`
	}{
		Limit: z.limit,
	}
	sf = make([]*mo_sharedfolder.SharedFolder, 0)
	req := z.ctx.List("sharing/list_folders").
		Continue("sharing/list_folders/continue").
		Param(p).
		UseHasMore(false).
		ResultTag("entries").
		OnEntry(func(entry api_list.ListEntry) error {
			f := &mo_sharedfolder.SharedFolder{}
			if err := entry.Model(f); err != nil {
				return err
			}
			sf = append(sf, f)
			return nil
		})
	if err := req.Call(); err != nil {
		return nil, err
	}
	return sf, err
}