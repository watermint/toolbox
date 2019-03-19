package sv_sharedlink

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/model/dbx_api"
	"time"
)

type SharedLink interface {
	List() (links []mo_sharedlink.SharedLink, err error)
	ListByPath(path mo_path.Path) (links []mo_sharedlink.SharedLink, err error)
	Delete(link mo_sharedlink.SharedLink) (err error)
	Create(path mo_path.Path, opts ...CreateOptions) (link mo_sharedlink.SharedLink, err error)
}

type createOption struct {
	visibility string
	password   string
	expires    string
}

type CreateOptions func(opt *createOption) *createOption

func Public() CreateOptions {
	return func(opt *createOption) *createOption {
		opt.visibility = "public"
		return opt
	}
}
func TeamOnly() CreateOptions {
	return func(opt *createOption) *createOption {
		opt.visibility = "team_only"
		return opt
	}
}
func Password(password string) CreateOptions {
	return func(opt *createOption) *createOption {
		opt.visibility = "password"
		opt.password = password
		return opt
	}
}
func Expires(at time.Time) CreateOptions {
	return func(opt *createOption) *createOption {
		opt.expires = dbx_api.RebaseTimeForAPI(at).Format(dbx_api.DateTimeFormat)
		return opt
	}
}

func New(ctx api_context.Context) SharedLink {
	return &sharedLinkImpl{
		ctx: ctx,
	}
}

type sharedLinkImpl struct {
	ctx api_context.Context
}

func (z *sharedLinkImpl) list(path string) (links []mo_sharedlink.SharedLink, err error) {
	p := struct {
		Path string `json:"path,omitempty"`
	}{
		Path: path,
	}

	links = make([]mo_sharedlink.SharedLink, 0)
	req := z.ctx.List("sharing/list_shared_links").
		Continue("sharing/list_shared_links").
		Param(p).
		UseHasMore(true).
		ResultTag("links").
		OnEntry(func(entry api_list.ListEntry) error {
			link := &mo_sharedlink.Metadata{}
			if err := entry.Model(link); err != nil {
				return err
			}
			links = append(links, link)
			return nil
		})
	if err := req.Call(); err != nil {
		return nil, err
	}
	return links, nil
}

func (z *sharedLinkImpl) List() (links []mo_sharedlink.SharedLink, err error) {
	return z.list("")
}

func (z *sharedLinkImpl) ListByPath(path mo_path.Path) (links []mo_sharedlink.SharedLink, err error) {
	return z.list(path.Path())
}

func (z *sharedLinkImpl) Delete(link mo_sharedlink.SharedLink) (err error) {
	p := struct {
		Url string `json:"url"`
	}{
		Url: link.LinkUrl(),
	}

	_, err = z.ctx.Request("sharing/revoke_shared_link").Param(p).Call()
	return err
}

func (z *sharedLinkImpl) Create(path mo_path.Path, opts ...CreateOptions) (link mo_sharedlink.SharedLink, err error) {
	opt := &createOption{}
	for _, o := range opts {
		o(opt)
	}
	type Settings struct {
		RequestedVisibility string `json:"requested_visibility,omitempty"`
		LinkPassword        string `json:"link_password,omitempty"`
		Expires             string `json:"expires,omitempty"`
	}
	p := struct {
		Path     string   `json:"path"`
		Settings Settings `json:"settings"`
	}{
		Path: path.Path(),
		Settings: Settings{
			RequestedVisibility: opt.visibility,
			LinkPassword:        opt.password,
			Expires:             opt.expires,
		},
	}

	link = &mo_sharedlink.Metadata{}
	res, err := z.ctx.Request("sharing/create_shared_link_with_settings").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err := res.Model(link); err != nil {
		return nil, err
	}
	return link, nil
}
