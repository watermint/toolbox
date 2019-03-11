package rp_file

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
)

type Repository interface {
}

func New(dc api_context.Context) Repository {
	r := &fileRepository{
		dc: dc,
	}
	return r
}

type fileRepository struct {
	dc                              api_context.Context
	recursive                       bool
	includeMediaInfo                bool
	includeDeleted                  bool
	includeHasExplicitSharedMembers bool
}

func (z *fileRepository) Delete(path mo_path.Path) (entry mo_file.Entry, err error) {
	panic("implement me")
}

func (z *fileRepository) List(path mo_path.Path) (entries []mo_file.Entry, err error) {
	entries = make([]mo_file.Entry, 0)
	p := struct {
		Path                            string `json:"path"`
		Recursive                       bool   `json:"recursive,omitempty"`
		IncludeMediaInfo                bool   `json:"include_media_info,omitempty"`
		IncludeDeleted                  bool   `json:"include_deleted,omitempty"`
		IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members,omitempty"`
	}{
		Path:                            path.Path(),
		Recursive:                       z.recursive,
		IncludeMediaInfo:                z.includeMediaInfo,
		IncludeDeleted:                  z.includeDeleted,
		IncludeHasExplicitSharedMembers: z.includeHasExplicitSharedMembers,
	}

	req := z.dc.List("files/list_folder").
		Continue("files/list_folder/continue").
		Param(p).
		UseHasMore(true).
		ResultTag("entries").
		OnEntry(func(entry api_list.ListEntry) error {
			e := &mo_file.Metadata{}
			if err := entry.Model(e); err != nil {
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

func (z *fileRepository) Resolve(path mo_path.Path) (entry mo_file.Entry, err error) {
	panic("implement me")
}
