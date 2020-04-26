package sv_file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/infra/api/api_request"
	"go.uber.org/zap"
)

type Files interface {
	Resolve(path mo_path.DropboxPath) (entry mo_file.Entry, err error)
	List(path mo_path.DropboxPath, opts ...ListOpt) (entries []mo_file.Entry, err error)
	ListChunked(path mo_path.DropboxPath, onEntry func(entry mo_file.Entry), opts ...ListOpt) error

	Remove(path mo_path.DropboxPath, opts ...RemoveOpt) (entry mo_file.Entry, err error)
	Poll(path mo_path.DropboxPath, onEntry func(entry mo_file.Entry), opts ...ListOpt) error
	Search(query string, opts ...SearchOpt) (matches []*mo_file.Match, err error)
}

type ListOpt func(opt *ListOpts) *ListOpts
type ListOpts struct {
	Recursive                       bool
	IncludeMediaInfo                bool
	IncludeDeleted                  bool
	IncludeHasExplicitSharedMembers bool
}

func Recursive() ListOpt {
	return func(opt *ListOpts) *ListOpts {
		opt.Recursive = true
		return opt
	}
}
func IncludeMediaInfo() ListOpt {
	return func(opt *ListOpts) *ListOpts {
		opt.IncludeMediaInfo = true
		return opt
	}
}
func IncludeDeleted() ListOpt {
	return func(opt *ListOpts) *ListOpts {
		opt.IncludeDeleted = true
		return opt
	}
}
func IncludeHasExplicitSharedMembers() ListOpt {
	return func(opt *ListOpts) *ListOpts {
		opt.IncludeHasExplicitSharedMembers = true
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

func NewFiles(ctx dbx_context.Context) Files {
	return &filesImpl{
		ctx: ctx,
	}
}

func newFilesTest(ctx dbx_context.Context) Files {
	return &filesImpl{
		ctx: ctx,
	}
}

type SearchOpt func(opt *searchOpts) *searchOpts
type searchOpts struct {
	path              string
	maxResults        *int
	fileStatus        string
	fileNameOnly      bool
	fileExtension     string
	fileCategories    string
	includeHighlights bool
}

func SearchPath(path mo_path.DropboxPath) SearchOpt {
	return func(opt *searchOpts) *searchOpts {
		opt.path = path.Path()
		return opt
	}
}
func SearchMaxResults(maxResults int) SearchOpt {
	return func(opt *searchOpts) *searchOpts {
		opt.maxResults = &maxResults
		return opt
	}
}
func SearchFileDeleted() SearchOpt {
	return func(opt *searchOpts) *searchOpts {
		opt.fileStatus = "deleted"
		return opt
	}
}
func SearchFileNameOnly() SearchOpt {
	return func(opt *searchOpts) *searchOpts {
		opt.fileNameOnly = true
		return opt
	}
}
func SearchFileExtension(ext string) SearchOpt {
	return func(opt *searchOpts) *searchOpts {
		opt.fileExtension = ext
		return opt
	}
}
func SearchCategories(cat string) SearchOpt {
	return func(opt *searchOpts) *searchOpts {
		opt.fileCategories = cat
		return opt
	}
}
func SearchIncludeHighlights() SearchOpt {
	return func(opt *searchOpts) *searchOpts {
		opt.includeHighlights = true
		return opt
	}
}

type filesImpl struct {
	ctx   dbx_context.Context
	limit int
}

func (z *filesImpl) Search(query string, opts ...SearchOpt) (matches []*mo_file.Match, err error) {
	so := &searchOpts{}
	for _, o := range opts {
		o(so)
	}
	type SO struct {
		MaxResults     *int   `json:"max_results,omitempty"`
		FileStatus     string `json:"file_status,omitempty"`
		FilenameOnly   bool   `json:"filename_only,omitempty"`
		FileExtensions string `json:"file_extensions,omitempty"`
		FileCategories string `json:"file_categories,omitempty"`
	}
	sos := &SO{
		MaxResults:     so.maxResults,
		FileStatus:     so.fileStatus,
		FilenameOnly:   so.fileNameOnly,
		FileExtensions: so.fileExtension,
		FileCategories: so.fileCategories,
	}
	p := struct {
		Query             string `json:"query"`
		Options           *SO    `json:"options"`
		IncludeHighlights bool   `json:"include_highlights,omitempty"`
	}{
		Query:             query,
		Options:           sos,
		IncludeHighlights: so.includeHighlights,
	}

	matches = make([]*mo_file.Match, 0)

	res := z.ctx.List("files/search_v2", api_request.Param(p)).Call(
		dbx_list.Continue("files/search/continue_v2"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("matches"),
		dbx_list.OnEntry(func(entry tjson.Json) error {
			e := &mo_file.Match{}
			if err := entry.Model(e); err != nil {
				z.ctx.Log().Error("invalid", zap.Error(err))
				return err
			}
			matches = append(matches, e)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return matches, nil
}

func (z *filesImpl) Poll(path mo_path.DropboxPath, onEntry func(entry mo_file.Entry), opts ...ListOpt) error {
	lo := &ListOpts{}
	for _, o := range opts {
		o(lo)
	}

	p := struct {
		Path                            string `json:"path"`
		Recursive                       bool   `json:"recursive,omitempty"`
		IncludeMediaInfo                bool   `json:"include_media_info,omitempty"`
		IncludeDeleted                  bool   `json:"include_deleted,omitempty"`
		IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members,omitempty"`
		Limit                           int    `json:"limit,omitempty"`
	}{
		Path:                            path.Path(),
		Recursive:                       lo.Recursive,
		IncludeMediaInfo:                lo.IncludeMediaInfo,
		IncludeDeleted:                  lo.IncludeDeleted,
		IncludeHasExplicitSharedMembers: lo.IncludeHasExplicitSharedMembers,
	}

	type Cursor struct {
		Cursor string `path:"cursor" json:"cursor"`
	}
	type LongPoll struct {
		Changes bool `path:"changes"  json:"changes"`
	}

	res := z.ctx.Post("files/list_folder/get_latest_cursor", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return err
	}
	cursor := &Cursor{}
	if err := res.Success().Json().Model(cursor); err != nil {
		return err
	}

	noAuthCtx := z.ctx.NoAuth()
	for {
		res := noAuthCtx.Notify("files/list_folder/longpoll", api_request.Param(cursor))
		if err, fail := res.Failure(); fail {
			return err
		}
		changes := &LongPoll{}
		if err := res.Success().Json().Model(changes); err != nil {
			return err
		}
		if changes.Changes {
			res := z.ctx.List("files/list_folder/continue", api_request.Param(cursor)).Call(
				dbx_list.Continue("files/list_folder/continue"),
				dbx_list.UseHasMore(),
				dbx_list.ResultTag("entries"),
				dbx_list.OnEntry(func(entry tjson.Json) error {
					e := &mo_file.Metadata{}
					if err := entry.Model(e); err != nil {
						z.ctx.Log().Error("invalid", zap.Error(err), zap.ByteString("entry", entry.Raw()))
						return err
					}
					onEntry(e)
					return nil
				}),
				dbx_list.OnLastCursor(func(c string) {
					cursor.Cursor = c
				}),
			)
			if err, fail := res.Failure(); fail {
				return err
			}
		}
	}
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
	res := z.ctx.Post("files/get_metadata", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	entry = &mo_file.Metadata{}
	err = res.Success().Json().Model(entry)
	return
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
	lo := &ListOpts{}
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
		Recursive:                       lo.Recursive,
		IncludeMediaInfo:                lo.IncludeMediaInfo,
		IncludeDeleted:                  lo.IncludeDeleted,
		IncludeHasExplicitSharedMembers: lo.IncludeHasExplicitSharedMembers,
	}

	res := z.ctx.List("files/list_folder", api_request.Param(p)).Call(
		dbx_list.Continue("files/list_folder/continue"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("entries"),
		dbx_list.OnEntry(func(entry tjson.Json) error {
			e := &mo_file.Metadata{}
			if err := entry.Model(e); err != nil {
				z.ctx.Log().Error("invalid", zap.Error(err), zap.ByteString("entry", entry.Raw()))
				return err
			}
			onEntry(e)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
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

	res := z.ctx.Post("files/delete_v2", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	entry = &mo_file.Metadata{}
	err = res.Success().Json().Model(entry)
	return
}
