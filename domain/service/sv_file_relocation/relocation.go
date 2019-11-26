package sv_file_relocation

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
)

type Relocation interface {
	// options: allow_shared_folder, allow_ownership_transfer, auto_rename
	Copy(from, to mo_path.Path) (entry mo_file.Entry, err error)

	// options: allow_shared_folder, allow_ownership_transfer, auto_rename
	Move(from, to mo_path.Path) (entry mo_file.Entry, err error)
}

type RelocationOption interface {
	AllowSharedFolder(allow bool)
	AllowOwnershipTransfer(allow bool)
	AutoRename(auto bool)
}

type Option func(option RelocationOption)

func New(ctx api_context.Context, options ...Option) Relocation {
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
	ctx                    api_context.Context
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

func (z *implRelocation) relocParam(from, to mo_path.Path) interface{} {
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

func (z *implRelocation) Copy(from, to mo_path.Path) (entry mo_file.Entry, err error) {
	p := z.relocParam(from, to)
	entry = &mo_file.Metadata{}
	res, err := z.ctx.Rpc("files/copy_v2").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err = res.ModelWithPath(entry, "metadata"); err != nil {
		return nil, err
	}
	return entry, nil
}

func (z *implRelocation) Move(from, to mo_path.Path) (entry mo_file.Entry, err error) {
	p := z.relocParam(from, to)
	entry = &mo_file.Metadata{}
	res, err := z.ctx.Rpc("files/move_v2").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err = res.ModelWithPath(entry, "metadata"); err != nil {
		return nil, err
	}
	return entry, nil
}
