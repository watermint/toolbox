package sv_sharedlink

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_util"
	"time"
)

type SharedLink interface {
	List() (links []mo_sharedlink.SharedLink, err error)
	ListByPath(path mo_path.DropboxPath) (links []mo_sharedlink.SharedLink, err error)
	Remove(link mo_sharedlink.SharedLink) (err error)
	Create(path mo_path.DropboxPath, opts ...LinkOpt) (link mo_sharedlink.SharedLink, err error)
	Update(link mo_sharedlink.SharedLink, opts ...LinkOpt) (updated mo_sharedlink.SharedLink, err error)
}

type linkOptions struct {
	visibility       string
	password         string
	expires          string
	removeExpiration bool
}

type LinkOpt func(opt *linkOptions) *linkOptions

func Public() LinkOpt {
	return func(opt *linkOptions) *linkOptions {
		opt.visibility = "public"
		return opt
	}
}
func TeamOnly() LinkOpt {
	return func(opt *linkOptions) *linkOptions {
		opt.visibility = "team_only"
		return opt
	}
}
func Password(password string) LinkOpt {
	return func(opt *linkOptions) *linkOptions {
		opt.visibility = "password"
		opt.password = password
		return opt
	}
}
func Expires(at time.Time) LinkOpt {
	return func(opt *linkOptions) *linkOptions {
		opt.expires = api_util.RebaseAsString(at)
		return opt
	}
}
func RemoveExpiration() LinkOpt {
	return func(opt *linkOptions) *linkOptions {
		opt.removeExpiration = true
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

func (z *sharedLinkImpl) Update(link mo_sharedlink.SharedLink, opts ...LinkOpt) (updated mo_sharedlink.SharedLink, err error) {
	opt := &linkOptions{}
	for _, o := range opts {
		o(opt)
	}
	type S struct {
		RequestedVisibility string `json:"requested_visibility,omitempty"`
		LinkPassword        string `json:"link_password,omitempty"`
		Expires             string `json:"expires,omitempty"`
	}
	p := struct {
		Url              string `json:"url"`
		Settings         S      `json:"settings"`
		RemoveExpiration bool   `json:"remove_expiration,omitempty"`
	}{
		Url: link.LinkUrl(),
		Settings: S{
			RequestedVisibility: opt.visibility,
			LinkPassword:        opt.password,
			Expires:             opt.expires,
		},
		RemoveExpiration: opt.removeExpiration,
	}

	link = &mo_sharedlink.Metadata{}
	res, err := z.ctx.Rpc("sharing/modify_shared_link_settings").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err := res.Model(link); err != nil {
		return nil, err
	}
	return link, nil
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

func (z *sharedLinkImpl) ListByPath(path mo_path.DropboxPath) (links []mo_sharedlink.SharedLink, err error) {
	return z.list(path.Path())
}

func (z *sharedLinkImpl) Remove(link mo_sharedlink.SharedLink) (err error) {
	p := struct {
		Url string `json:"url"`
	}{
		Url: link.LinkUrl(),
	}

	_, err = z.ctx.Rpc("sharing/revoke_shared_link").Param(p).Call()
	return err
}

func (z *sharedLinkImpl) Create(path mo_path.DropboxPath, opts ...LinkOpt) (link mo_sharedlink.SharedLink, err error) {
	opt := &linkOptions{}
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
	res, err := z.ctx.Rpc("sharing/create_shared_link_with_settings").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err := res.Model(link); err != nil {
		return nil, err
	}
	return link, nil
}
