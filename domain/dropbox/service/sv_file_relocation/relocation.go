package sv_file_relocation

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/api/api_request"
)

type Relocation interface {
	// options: allow_shared_folder, allow_ownership_transfer, auto_rename
	Copy(from, to mo_path.DropboxPath) (entry mo_file.Entry, err error)

	// options: allow_shared_folder, allow_ownership_transfer, auto_rename
	Move(from, to mo_path.DropboxPath) (entry mo_file.Entry, err error)
}

type RelocationOption interface {
	AllowSharedFolder(allow bool)
	AllowOwnershipTransfer(allow bool)
	AutoRename(auto bool)
}

type Option func(option RelocationOption)

func New(ctx dbx_client.Client, options ...Option) Relocation {
	r := &implRelocation{
		ctx: ctx,
	}
	for _, op := range options {
		op(r)
	}
	return r
}

func AllowSharedFolder(allow bool) Option {
	return func(op RelocationOption) {
		op.AllowSharedFolder(allow)
	}
}

func AllowOwnershipTransfer(allow bool) Option {
	return func(op RelocationOption) {
		op.AllowOwnershipTransfer(allow)
	}
}

func AutoRename(auto bool) Option {
	return func(op RelocationOption) {
		op.AutoRename(auto)
	}
}

type implRelocation struct {
	ctx                    dbx_client.Client
	allowSharedFolder      bool
	allowOwnershipTransfer bool
	autoRename             bool
}

func (z *implRelocation) AllowSharedFolder(allow bool) {
	z.allowSharedFolder = allow
}

func (z *implRelocation) AllowOwnershipTransfer(allow bool) {
	z.allowOwnershipTransfer = allow
}

func (z *implRelocation) AutoRename(auto bool) {
	z.autoRename = auto
}

func (z *implRelocation) relocParam(from, to mo_path.DropboxPath) interface{} {
	p := struct {
		FromPath               string `json:"from_path"`
		ToPath                 string `json:"to_path"`
		AllowSharedFolder      bool   `json:"allow_shared_folder,omitempty"`
		AllowOwnershipTransfer bool   `json:"allow_ownership_transfer,omitempty"`
		AutoRename             bool   `json:"auto_rename,omitempty"`
	}{
		FromPath:               from.Path(),
		ToPath:                 to.Path(),
		AllowSharedFolder:      z.allowSharedFolder,
		AllowOwnershipTransfer: z.allowOwnershipTransfer,
		AutoRename:             z.autoRename,
	}
	return p
}

func (z *implRelocation) Copy(from, to mo_path.DropboxPath) (entry mo_file.Entry, err error) {
	p := z.relocParam(from, to)
	res := z.ctx.Post("files/copy_v2", api_request.Param(p))
	if err, f := res.Failure(); f {
		return nil, err
	}
	entry = &mo_file.Metadata{}
	err = res.Success().Json().FindModel("metadata", entry)
	return
}

func (z *implRelocation) Move(from, to mo_path.DropboxPath) (entry mo_file.Entry, err error) {
	p := z.relocParam(from, to)
	res := z.ctx.Post("files/move_v2", api_request.Param(p))
	if err, f := res.Failure(); f {
		return nil, err
	}
	entry = &mo_file.Metadata{}
	err = res.Success().Json().FindModel("metadata", entry)
	return
}
