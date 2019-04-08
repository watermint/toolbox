package sv_file

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"go.uber.org/zap"
)

type Files interface {
	Resolve(path mo_path.Path) (entry mo_file.Entry, err error)
	List(path mo_path.Path) (entries []mo_file.Entry, err error)

	Remove(path mo_path.Path, opts ...RemoveOpt) (entry mo_file.Entry, err error)
}

type RemoveOpt func(opt *removeOpts) *removeOpts
type removeOpts struct {
	permanently bool
	revision    string
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
	ctx                             api_context.Context
	recursive                       bool
	includeMediaInfo                bool
	includeDeleted                  bool
	includeHasExplicitSharedMembers bool
	limit                           int
}

func (z *filesImpl) Resolve(path mo_path.Path) (entry mo_file.Entry, err error) {
	p := struct {
		Path                            string `json:"path"`
		IncludeMediaInfo                bool   `json:"include_media_info,omitempty"`
		IncludeDeleted                  bool   `json:"include_deleted,omitempty"`
		IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members,omitempty"`
	}{
		Path:                            path.Path(),
		IncludeMediaInfo:                z.includeMediaInfo,
		IncludeDeleted:                  z.includeDeleted,
		IncludeHasExplicitSharedMembers: z.includeHasExplicitSharedMembers,
	}
	entry = &mo_file.Metadata{}
	res, err := z.ctx.Request("files/get_metadata").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err := res.Model(entry); err != nil {
		return nil, err
	}
	return entry, nil
}

func (z *filesImpl) List(path mo_path.Path) (entries []mo_file.Entry, err error) {
	entries = make([]mo_file.Entry, 0)
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
		Recursive:                       z.recursive,
		IncludeMediaInfo:                z.includeMediaInfo,
		IncludeDeleted:                  z.includeDeleted,
		IncludeHasExplicitSharedMembers: z.includeHasExplicitSharedMembers,
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
			entries = append(entries, e)
			return nil
		})
	if err := req.Call(); err != nil {
		return nil, err
	}
	return entries, nil
}

func (z *filesImpl) Remove(path mo_path.Path, opts ...RemoveOpt) (entry mo_file.Entry, err error) {
	panic("implement me")
}
