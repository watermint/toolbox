package api_auth_impl

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/experimental/app_control"
)

func Auth(ec *app.ExecContext, opts ...AuthOpt) (ctx api_context.Context, err error) {
	ao := &authOpts{
		tokenType: api_auth.DropboxTokenNoAuth,
		peerName:  ec.DefaultPeerName(),
	}
	for _, o := range opts {
		o(ao)
	}
	a := New(ec, opts...)
	return a.Auth(ao.tokenType)
}

func NewKc(control app_control.Control, opts ...AuthOpt) api_auth.Console {
	ao := &authOpts{
		tokenType: api_auth.DropboxTokenNoAuth,
		peerName:  "default",
	}
	for _, o := range opts {
		o(ao)
	}
	ua := &KitchenAuth{
		control: control,
	}
	ua.init()
	ca := &KcCachedAuth{
		peerName: ao.peerName,
		control:  control,
		auth:     ua,
	}
	ca.init()
	return ca
}

func New(ec *app.ExecContext, opts ...AuthOpt) api_auth.Console {
	ao := &authOpts{
		tokenType: api_auth.DropboxTokenNoAuth,
		peerName:  ec.DefaultPeerName(),
	}
	for _, o := range opts {
		o(ao)
	}
	ua := &ECAuth{
		ec: ec,
	}
	ua.init()
	ca := &EcCachedAuth{
		peerName: ao.peerName,
		ec:       ec,
		auth:     ua,
	}
	ca.init()
	return ca
}

func IsCacheAvailable(ec *app.ExecContext, peerName string) bool {
	ca := &EcCachedAuth{
		peerName: peerName,
		ec:       ec,
	}
	ca.init()
	return len(ca.tokens) > 4
}

type AuthOpt func(opt *authOpts) *authOpts
type authOpts struct {
	peerName  string
	tokenType string
}

func PeerName(name string) AuthOpt {
	return func(opt *authOpts) *authOpts {
		opt.peerName = name
		return opt
	}
}
func Full() AuthOpt {
	return func(opt *authOpts) *authOpts {
		opt.tokenType = api_auth.DropboxTokenFull
		return opt
	}
}
func BusinessFile() AuthOpt {
	return func(opt *authOpts) *authOpts {
		opt.tokenType = api_auth.DropboxTokenBusinessFile
		return opt
	}
}
func BusinessManagement() AuthOpt {
	return func(opt *authOpts) *authOpts {
		opt.tokenType = api_auth.DropboxTokenBusinessManagement
		return opt
	}
}
func BusinessInfo() AuthOpt {
	return func(opt *authOpts) *authOpts {
		opt.tokenType = api_auth.DropboxTokenBusinessInfo
		return opt
	}
}
func BusinessAudit() AuthOpt {
	return func(opt *authOpts) *authOpts {
		opt.tokenType = api_auth.DropboxTokenBusinessAudit
		return opt
	}
}
