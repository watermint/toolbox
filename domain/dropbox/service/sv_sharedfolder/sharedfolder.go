package sv_sharedfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type SharedFolder interface {
	Create(path mo_path.DropboxPath, opts ...CreateOpt) (sf *mo_sharedfolder.SharedFolder, err error)
	Remove(sf *mo_sharedfolder.SharedFolder, opts ...DeleteOpt) (err error)
	List() (sf []*mo_sharedfolder.SharedFolder, err error)
	Leave(sf *mo_sharedfolder.SharedFolder, opts ...DeleteOpt) (err error)
	Resolve(sharedFolderId string) (sf *mo_sharedfolder.SharedFolder, err error)
	Transfer(sf *mo_sharedfolder.SharedFolder, to TransferTo) (err error)
	UpdatePolicy(sharedFolderId string, opts ...PolicyOpt) (sf *mo_sharedfolder.SharedFolder, err error)
	UpdateInheritance(sharedFolderId string, setting string) (sf *mo_sharedfolder.SharedFolder, err error)
}

func New(ctx dbx_context.Context) SharedFolder {
	return &sharedFolderImpl{
		ctx: ctx,
	}
}

const (
	AccessInheritanceInherit   = "inherit"
	AccessInheritanceNoInherit = "no_inherit"
)

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
func ToAccountId(accountId string) TransferTo {
	return func(to *transferTo) *transferTo {
		to.dropboxId = accountId
		return to
	}
}
func ToTeamMemberId(teamMemberId string) TransferTo {
	return func(to *transferTo) *transferTo {
		to.dropboxId = teamMemberId
		return to
	}
}

type PolicyOpt func(opt *policyOpts) *policyOpts
type policyOpts struct {
	SharedFolderId   string `json:"shared_folder_id"`
	MemberPolicy     string `json:"member_policy"`
	AclUpdatePolicy  string `json:"acl_update_policy"`
	SharedLinkPolicy string `json:"shared_link_policy"`
}

func MemberPolicy(policy string) PolicyOpt {
	return func(opt *policyOpts) *policyOpts {
		opt.MemberPolicy = policy
		return opt
	}
}
func AclUpdatePolicy(policy string) PolicyOpt {
	return func(opt *policyOpts) *policyOpts {
		opt.AclUpdatePolicy = policy
		return opt
	}
}
func SharedLinkPolicy(policy string) PolicyOpt {
	return func(opt *policyOpts) *policyOpts {
		opt.SharedLinkPolicy = policy
		return opt
	}
}

type createOpts struct {
}

type CreateOpt func(opt *createOpts) *createOpts

type deleteOpts struct {
	leaveACopy bool
}
type DeleteOpt func(opt *deleteOpts) *deleteOpts

func LeaveACopy() DeleteOpt {
	return func(opt *deleteOpts) *deleteOpts {
		opt.leaveACopy = true
		return opt
	}
}

type sharedFolderImpl struct {
	ctx   dbx_context.Context
	limit int
}

func (z *sharedFolderImpl) UpdateInheritance(sharedFolderId string, setting string) (sf *mo_sharedfolder.SharedFolder, err error) {
	p := struct {
		SharedFolderId    string `json:"shared_folder_id"`
		AccessInheritance string `json:"access_inheritance"`
	}{
		SharedFolderId:    sharedFolderId,
		AccessInheritance: setting,
	}

	res := z.ctx.Post("sharing/set_access_inheritance", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	sf = &mo_sharedfolder.SharedFolder{}
	err = res.Success().Json().Model(sf)
	return
}

func (z *sharedFolderImpl) UpdatePolicy(sharedFolderId string, opts ...PolicyOpt) (sf *mo_sharedfolder.SharedFolder, err error) {
	po := &policyOpts{}
	for _, o := range opts {
		o(po)
	}
	po.SharedFolderId = sharedFolderId
	res := z.ctx.Post("sharing/update_folder_policy", api_request.Param(po))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	sf = &mo_sharedfolder.SharedFolder{}
	err = res.Success().Json().Model(sf)
	return
}

func (z *sharedFolderImpl) Transfer(sf *mo_sharedfolder.SharedFolder, to TransferTo) (err error) {
	too := &transferTo{}
	to(too)

	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
		ToDropboxId    string `json:"to_dropbox_id,omitempty"`
	}{
		SharedFolderId: sf.SharedFolderId,
		ToDropboxId:    too.dropboxId,
	}

	res := z.ctx.Post("sharing/transfer_folder", api_request.Param(p))
	if err, fail := res.Failure(); fail {
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

	res := z.ctx.Post("sharing/get_folder_metadata", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	sf = &mo_sharedfolder.SharedFolder{}
	err = res.Success().Json().Model(sf)
	return
}

func (z *sharedFolderImpl) Leave(sf *mo_sharedfolder.SharedFolder, opts ...DeleteOpt) (err error) {
	do := &deleteOpts{}
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
	res := z.ctx.Async("sharing/relinquish_folder_membership", api_request.Param(p)).Call(
		dbx_async.Status("sharing/check_share_job_status"))
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}

func (z *sharedFolderImpl) Create(path mo_path.DropboxPath, opts ...CreateOpt) (sf *mo_sharedfolder.SharedFolder, err error) {
	co := &createOpts{}
	for _, o := range opts {
		o(co)
	}

	p := struct {
		Path string `json:"path"`
	}{
		Path: path.Path(),
	}

	res := z.ctx.Async("sharing/share_folder", api_request.Param(p)).Call(
		dbx_async.Status("sharing/check_share_job_status"))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	sf = &mo_sharedfolder.SharedFolder{}
	err = res.Success().Json().Model(sf)
	return
}

func (z *sharedFolderImpl) Remove(sf *mo_sharedfolder.SharedFolder, opts ...DeleteOpt) (err error) {
	do := &deleteOpts{}
	for _, o := range opts {
		o(do)
	}
	sharedFolderId := sf.SharedFolderId
	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
		LeaveACopy     bool   `json:"leave_a_copy,omitempty"`
	}{
		SharedFolderId: sharedFolderId,
		LeaveACopy:     do.leaveACopy,
	}
	sf = &mo_sharedfolder.SharedFolder{}
	res := z.ctx.Async("sharing/unshare_folder", api_request.Param(p)).Call(
		dbx_async.Status("sharing/check_job_status"))
	if err, fail := res.Failure(); fail {
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
	res := z.ctx.List("sharing/list_folders", api_request.Param(p)).Call(
		dbx_list.Continue("sharing/list_folders/continue"),
		dbx_list.ResultTag("entries"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			f := &mo_sharedfolder.SharedFolder{}
			if err := entry.Model(f); err != nil {
				return err
			}
			sf = append(sf, f)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return
}
