package sv_sharedlink_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_url"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_request"
	"strings"
)

type File interface {
	List(url mo_url.Url, path mo_path.DropboxPath, nEntry func(entry mo_file.Entry), opt ...ListOpt) error
	ListRecursive(url mo_url.Url, nEntry func(entry mo_file.Entry), opt ...ListOpt) error
}

type ListOpt func(opt *ListOpts) *ListOpts
type ListOpts struct {
	IncludeDeleted                  bool
	IncludeHasExplicitSharedMembers bool
	Password                        string
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
func Password(password string) ListOpt {
	return func(opt *ListOpts) *ListOpts {
		opt.Password = password
		return opt
	}
}

func New(ctx dbx_context.Context) File {
	return &fileImpl{ctx: ctx}
}

type fileImpl struct {
	ctx dbx_context.Context
}

func (z *fileImpl) ListRecursive(url mo_url.Url, nEntry func(entry mo_file.Entry), opts ...ListOpt) error {
	lo := &ListOpts{}
	for _, o := range opts {
		o(lo)
	}
	var ls func(entry0 mo_file.Entry, rel string) error
	ls = func(entry0 mo_file.Entry, rel string) error {
		c := entry0.Concrete()
		c.PathDisplay = rel
		if !strings.HasPrefix(c.PathDisplay, "/") {
			c.PathDisplay = "/" + c.PathDisplay
		}
		c.PathLower = strings.ToLower(c.PathDisplay) // TODO: i18n
		r := make(map[string]interface{})
		if err0 := json.Unmarshal(c.Raw, &r); err0 == nil {
			r["path_display"] = c.PathDisplay
			r["path_lower"] = c.PathLower
			if r0, err0 := json.Marshal(&r); err0 == nil {
				c.Raw = r0
			}
		}

		if f, ok := entry0.File(); ok {
			f.Raw = c.Raw
			f.EntryPathDisplay = c.PathDisplay
			f.EntryPathLower = c.PathLower
			nEntry(f)
			return nil
		}
		if f, ok := entry0.Deleted(); ok {
			f.Raw = c.Raw
			f.EntryPathDisplay = c.PathDisplay
			f.EntryPathLower = c.PathLower
			nEntry(f)
			return nil
		}
		if f, ok := entry0.Folder(); ok {
			f.Raw = c.Raw
			f.EntryPathDisplay = c.PathDisplay
			f.EntryPathLower = c.PathLower
			nEntry(f)
		}

		return z.List(url, mo_path.NewDropboxPath(rel), func(entry1 mo_file.Entry) {
			ls(entry1, rel+"/"+entry1.Name())
		}, opts...)
	}

	entry, err := sv_sharedlink.New(z.ctx).Resolve(url, lo.Password)
	if err != nil {
		return err
	}

	return ls(entry, "")
}

func (z *fileImpl) List(url mo_url.Url, path mo_path.DropboxPath, onEntry func(entry mo_file.Entry), opts ...ListOpt) error {
	lo := &ListOpts{}
	for _, o := range opts {
		o(lo)
	}

	type SL struct {
		Url      string `json:"url"`
		Password string `json:"password,omitempty"`
	}
	p := struct {
		Path                            string `json:"path"`
		SharedLink                      *SL    `json:"shared_link"`
		Recursive                       bool   `json:"recursive,omitempty"`
		IncludeDeleted                  bool   `json:"include_deleted,omitempty"`
		IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members,omitempty"`
		Limit                           int    `json:"limit,omitempty"`
	}{
		Path: path.Path(),
		SharedLink: &SL{
			Url:      url.Value(),
			Password: lo.Password,
		},
		IncludeDeleted:                  lo.IncludeDeleted,
		IncludeHasExplicitSharedMembers: lo.IncludeHasExplicitSharedMembers,
	}

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
