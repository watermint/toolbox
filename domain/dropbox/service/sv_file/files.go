package sv_file

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_request"
)

var (
	ErrorEntryNotFound = errors.New("entry not found")
)

type Files interface {
	Resolve(path mo_path.DropboxPath, opts ...ResolveOpt) (entry mo_file.Entry, err error)
	List(path mo_path.DropboxPath, opts ...ListOpt) (entries []mo_file.Entry, err error)
	ListEach(path mo_path.DropboxPath, onEntry func(entry mo_file.Entry), opts ...ListOpt) error

	Remove(path mo_path.DropboxPath, opts ...RemoveOpt) (entry mo_file.Entry, err error)
	PermDelete(path mo_path.DropboxPath) (err error)
	Poll(path mo_path.DropboxPath, onEntry func(entry mo_file.Entry), opts ...ListOpt) error
	Search(query string, opts ...SearchOpt) (matches []*mo_file.Match, err error)
	UploadLink(path mo_path.DropboxPath) (url string, err error)

	Lock(path mo_path.DropboxPath) (entry mo_file.Entry, err error)
	Unlock(path mo_path.DropboxPath) (entry mo_file.Entry, err error)
}

type ResolveOpt func(opt ResolveOpts) ResolveOpts

type ResolveOpts struct {
	Path                            string `json:"path"`
	IncludeDeleted                  bool   `json:"include_deleted,omitempty"`
	IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members,omitempty"`
}

func (z ResolveOpts) Apply(opts []ResolveOpt) ResolveOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

func ResolveIncludeDeleted(enabled bool) ResolveOpt {
	return func(opt ResolveOpts) ResolveOpts {
		opt.IncludeDeleted = enabled
		return opt
	}
}
func ResolveIncludeHasExplicitSharedMembers(enabled bool) ResolveOpt {
	return func(opt ResolveOpts) ResolveOpts {
		opt.IncludeDeleted = enabled
		return opt
	}
}

type ListOpt func(opt ListOpts) ListOpts
type ListOpts struct {
	Path                            string `json:"path"`
	Limit                           *int   `json:"limit,omitempty"`
	Recursive                       bool   `json:"recursive,omitempty"`
	IncludeDeleted                  bool   `json:"include_deleted,omitempty"`
	IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members,omitempty"`
	IncludeNonDownloadableFiles     bool   `json:"include_non_downloadable_files"` // should not omitempty, default is true
}

func MakeListOpts(path mo_path.DropboxPath, opts []ListOpt) ListOpts {
	o := ListOpts{
		Path:                            path.Path(),
		Limit:                           nil,
		Recursive:                       false,
		IncludeDeleted:                  false,
		IncludeHasExplicitSharedMembers: false,
		IncludeNonDownloadableFiles:     true,
	}
	for _, opt := range opts {
		o = opt(o)
	}
	return o
}

func Recursive(enabled bool) ListOpt {
	return func(opt ListOpts) ListOpts {
		opt.Recursive = enabled
		return opt
	}
}
func IncludeDeleted(enabled bool) ListOpt {
	return func(opt ListOpts) ListOpts {
		opt.IncludeDeleted = enabled
		return opt
	}
}
func IncludeHasExplicitSharedMembers(enabled bool) ListOpt {
	return func(opt ListOpts) ListOpts {
		opt.IncludeHasExplicitSharedMembers = enabled
		return opt
	}
}
func IncludeNonDownloadableFiles(enabled bool) ListOpt {
	return func(opt ListOpts) ListOpts {
		opt.IncludeNonDownloadableFiles = enabled
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
	fileCategories    []string
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
func SearchCategories(cat ...string) SearchOpt {
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

func (z *filesImpl) lockOrUnlock(path mo_path.DropboxPath, endpoint string) (entry mo_file.Entry, err error) {
	type PA struct {
		Path string `json:"path"`
	}
	type PB struct {
		Entries []PA `json:"entries"`
	}
	p := PB{
		[]PA{
			{path.Path()},
		},
	}
	res := z.ctx.Post(endpoint, api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}

	je, found := res.Success().Json().Find("entries.0")
	if !found {
		return nil, ErrorEntryNotFound
	}
	tag, found := je.FindString("\\.tag")
	if !found {
		return nil, ErrorEntryNotFound
	}
	if tag == "success" {
		entry = &mo_file.Metadata{}
		err = je.FindModel("metadata", entry)
		return entry, err
	} else {
		reason, found := je.FindString("failure.\\.tag")
		if !found {
			return nil, ErrorEntryNotFound
		}
		detail, found := je.FindString("failure." + reason + "\\.tag")
		if !found {
			return nil, errors.New(reason)
		}
		return nil, errors.New(reason + "/" + detail)
	}
}

func (z *filesImpl) Lock(path mo_path.DropboxPath) (entry mo_file.Entry, err error) {
	return z.lockOrUnlock(path, "files/lock_file_batch")
}

func (z *filesImpl) Unlock(path mo_path.DropboxPath) (entry mo_file.Entry, err error) {
	return z.lockOrUnlock(path, "files/unlock_file_batch")
}

func (z *filesImpl) UploadLink(path mo_path.DropboxPath) (url string, err error) {
	type CommitInfo struct {
		Path       string `json:"path"`
		AutoRename bool   `json:"autorename"`
	}
	p := struct {
		CommitInfo CommitInfo `json:"commit_info"`
	}{
		CommitInfo: CommitInfo{
			Path:       path.Path(),
			AutoRename: false,
		},
	}
	res := z.ctx.Post("files/get_temporary_upload_link", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return "", err
	}

	link, found := res.Success().Json().FindString("link")
	if found {
		return link, nil
	}
	return "", errors.New("invalid response")
}

func (z *filesImpl) PermDelete(path mo_path.DropboxPath) (err error) {
	p := struct {
		Path string `json:"path"`
	}{
		Path: path.Path(),
	}
	res := z.ctx.Post("files/permanently_delete", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}

func (z *filesImpl) Search(query string, opts ...SearchOpt) (matches []*mo_file.Match, err error) {
	so := &searchOpts{}
	for _, o := range opts {
		o(so)
	}
	type SO struct {
		MaxResults     *int     `json:"max_results,omitempty"`
		FileStatus     string   `json:"file_status,omitempty"`
		FilenameOnly   bool     `json:"filename_only,omitempty"`
		FileExtensions string   `json:"file_extensions,omitempty"`
		FileCategories []string `json:"file_categories,omitempty"`
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
		dbx_list.OnEntry(func(entry es_json.Json) error {
			e := &mo_file.Match{}
			if err := entry.Model(e); err != nil {
				z.ctx.Log().Error("invalid", esl.Error(err))
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
	p := MakeListOpts(path, opts)

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
				dbx_list.OnEntry(func(entry es_json.Json) error {
					e := &mo_file.Metadata{}
					if err := entry.Model(e); err != nil {
						z.ctx.Log().Error("invalid", esl.Error(err), esl.ByteString("entry", entry.Raw()))
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

func (z *filesImpl) Resolve(path mo_path.DropboxPath, opts ...ResolveOpt) (entry mo_file.Entry, err error) {
	p := ResolveOpts{Path: path.Path()}.Apply(opts)
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
	err = z.ListEach(path, func(entry mo_file.Entry) {
		entries = append(entries, entry)
	}, opts...)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (z *filesImpl) ListEach(path mo_path.DropboxPath, onEntry func(entry mo_file.Entry), opts ...ListOpt) error {
	p := MakeListOpts(path, opts)

	res := z.ctx.List("files/list_folder", api_request.Param(p)).Call(
		dbx_list.Continue("files/list_folder/continue"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("entries"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			e := &mo_file.Metadata{}
			if err := entry.Model(e); err != nil {
				z.ctx.Log().Error("invalid", esl.Error(err), esl.ByteString("entry", entry.Raw()))
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
