package sv_file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"go.uber.org/zap"
)

type Files interface {
	Resolve(path mo_path.DropboxPath) (entry mo_file.Entry, err error)
	List(path mo_path.DropboxPath, opts ...ListOpt) (entries []mo_file.Entry, err error)
	ListChunked(path mo_path.DropboxPath, onEntry func(entry mo_file.Entry), opts ...ListOpt) error

	Remove(path mo_path.DropboxPath, opts ...RemoveOpt) (entry mo_file.Entry, err error)
}

type ListOpt func(opt *listOpts) *listOpts
type listOpts struct {
	recursive                       bool
	includeMediaInfo                bool
	includeDeleted                  bool
	includeHasExplicitSharedMembers bool
}

func Recursive() ListOpt {
	return func(opt *listOpts) *listOpts {
		opt.recursive = true
		return opt
	}
}
func IncludeMediaInfo() ListOpt {
	return func(opt *listOpts) *listOpts {
		opt.includeMediaInfo = true
		return opt
	}
}
func IncludeDeleted() ListOpt {
	return func(opt *listOpts) *listOpts {
		opt.includeDeleted = true
		return opt
	}
}
func IncludeHasExplicitSharedMembers() ListOpt {
	return func(opt *listOpts) *listOpts {
		opt.includeHasExplicitSharedMembers = true
		return opt
	}
}

type RemoveOpt func(opt *removeOpts) *removeOpts
type removeOpts struct {
	revision string
}

func RemoveRevision(revision string) RemoveOpt {
	return func(opt *removeOpts) *removeOpts {
		opt.revision = revision
		return opt
	}
}

func NewFiles(ctx api_context.Context) Files {
	return &filesImpl{
		ctx: ctx,
	}
}

func newFilesTest(ctx api_context.Context) Files {
	return &filesImpl{
		ctx: ctx,
		//limit: 3,
	}
}

type filesImpl struct {
	ctx   api_context.Context
	limit int
}

func (z *filesImpl) Resolve(path mo_path.DropboxPath) (entry mo_file.Entry, err error) {
	p := struct {
		Path                            string `json:"path"`
		IncludeMediaInfo                bool   `json:"include_media_info,omitempty"`
		IncludeDeleted                  bool   `json:"include_deleted,omitempty"`
		IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members,omitempty"`
	}{
		Path:                            path.Path(),
		IncludeHasExplicitSharedMembers: true,
		IncludeMediaInfo:                true,
		IncludeDeleted:                  false,
	}
	entry = &mo_file.Metadata{}
	res, err := z.ctx.Rpc("files/get_metadata").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err := res.Model(entry); err != nil {
		return nil, err
	}
	return entry, nil
}

func (z *filesImpl) List(path mo_path.DropboxPath, opts ...ListOpt) (entries []mo_file.Entry, err error) {
	entries = make([]mo_file.Entry, 0)
	err = z.ListChunked(path, func(entry mo_file.Entry) {
		entries = append(entries, entry)
	}, opts...)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (z *filesImpl) ListChunked(path mo_path.DropboxPath, onEntry func(entry mo_file.Entry), opts ...ListOpt) error {
	lo := &listOpts{}
	for _, o := range opts {
		o(lo)
	}

	pp := path.Path()
	if pp == "/" {
		pp = ""
	}

	p := struct {
		Path                            string `json:"path"`
		Recursive                       bool   `json:"recursive,omitempty"`
		IncludeMediaInfo                bool   `json:"include_media_info,omitempty"`
		IncludeDeleted                  bool   `json:"include_deleted,omitempty"`
		IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members,omitempty"`
		Limit                           int    `json:"limit,omitempty"`
	}{
		Path:                            pp,
		Recursive:                       lo.recursive,
		IncludeMediaInfo:                lo.includeMediaInfo,
		IncludeDeleted:                  lo.includeDeleted,
		IncludeHasExplicitSharedMembers: lo.includeHasExplicitSharedMembers,
	}

	req := z.ctx.List("files/list_folder").
		Continue("files/list_folder/continue").
		Param(p).
		UseHasMore(true).
		ResultTag("entries").
		OnEntry(func(entry api_list.ListEntry) error {
			e := &mo_file.Metadata{}
			if err := entry.Model(e); err != nil {
				j, _ := entry.Json()
				z.ctx.Log().Error("invalid", zap.Error(err), zap.String("entry", j.Raw))
				return err
			}
			onEntry(e)
			return nil
		})
	return req.Call()
}

func (z *filesImpl) Remove(path mo_path.DropboxPath, opts ...RemoveOpt) (entry mo_file.Entry, err error) {
	opt := &removeOpts{}
	for _, o := range opts {
		o(opt)
	}

	p := struct {
		Path      string `json:"path"`
		ParentRev string `json:"parent_rev,omitempty"`
	}{
		Path:      path.Path(),
		ParentRev: opt.revision,
	}

	entry = &mo_file.Metadata{}
	res, err := z.ctx.Rpc("files/delete_v2").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err := res.Model(entry); err != nil {
		return nil, err
	}
	return entry, nil
}
